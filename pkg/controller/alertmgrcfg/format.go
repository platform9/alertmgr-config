package alertmgrcfg

import (
	"fmt"
	"os"

	monitoringv1alpha1 "github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1"
	"github.com/platform9/alertmgr-config/pkg/util"
)

const (
	suffixLen = 8
)

type format interface {
	formatAlert(amc *monitoringv1alpha1.AlertMgrCfg, acfg *alertConfig) error
}

func getFormatter(ftype string) (format, error) {
	var f format
	switch ftype {
	case "slack":
		f = slackconfig{}
	case "email":
		f = emailconfig{}
	default:
		return nil, os.ErrInvalid
	}

	return f, nil
}

func formatReceiver(amc *monitoringv1alpha1.AlertMgrCfg, acfg *alertConfig) error {

	var f format
	f, err := getFormatter(amc.Spec.Type)
	if err != nil {
		return err
	}

	err = f.formatAlert(amc, acfg)
	if err != nil {
		return err
	}

	return nil
}

func (f slackconfig) formatAlert(amc *monitoringv1alpha1.AlertMgrCfg, acfg *alertConfig) error {
	var url, channel, severity string
	for _, param := range amc.Spec.Params {
		switch param.Name {
		case "url":
			url = param.Value
		case "channel":
			channel = param.Value
		case "severity":
			severity = param.Value
		}
	}

	if url == "" {
		return os.ErrInvalid
	}

	if channel == "" {
		return os.ErrInvalid
	}

	if severity == "" {
		return os.ErrInvalid
	}

	if acfg.Route.Routes == nil {
		acfg.Route.Routes = []routes{}
	}
	receiverName := fmt.Sprintf("%s-%s", "slack", util.RandString(suffixLen))

	acfg.Route.Routes = append(acfg.Route.Routes, routes{
		Receiver: receiverName,
		MatchRe: map[string]string{
			"severity": severity,
		},
	})

	acfg.Receivers = append(acfg.Receivers, receiver{
		Name: receiverName,
		SlackConfigs: []slackconfig{
			slackconfig{
				ApiURL:  url,
				Channel: channel,
			},
		},
	})

	return nil
}

func (f emailconfig) formatAlert(amc *monitoringv1alpha1.AlertMgrCfg, acfg *alertConfig) error {
	var to, from, smarthost, severity string
	var auth_identity, auth_username, auth_password string
	for _, param := range amc.Spec.Params {
		switch param.Name {
		case "to":
			to = param.Value
		case "from":
			from = param.Value
		case "smarthost":
			smarthost = param.Value
		case "severity":
			severity = param.Value
		case "auth_identity":
			auth_identity = param.Value
		case "auth_username":
			auth_username = param.Value
		case "auth_password":
			auth_password = param.Value
		}
	}

	if to == "" {
		return os.ErrInvalid
	}

	if from == "" {
		return os.ErrInvalid
	}

	if smarthost == "" {
		return os.ErrInvalid
	}

	if severity == "" {
		return os.ErrInvalid
	}

	if auth_identity == "" {
		return os.ErrInvalid
	}

	if auth_username == "" {
		return os.ErrInvalid
	}

	if auth_password == "" {
		return os.ErrInvalid
	}

	if acfg.Route.Routes == nil {
		acfg.Route.Routes = []routes{}
	}

	receiverName := fmt.Sprintf("%s-%s", "email", util.RandString(suffixLen))
	acfg.Route.Routes = append(acfg.Route.Routes, routes{
		Receiver: receiverName,
		MatchRe: map[string]string{
			"severity": severity,
		},
	})

	acfg.Receivers = append(acfg.Receivers, receiver{
		Name: receiverName,
		EmailConfigs: []emailconfig{
			emailconfig{
				To:           to,
				From:         from,
				SmartHost:    smarthost,
				AuthUsername: auth_username,
				AuthIdentity: auth_identity,
				AuthPassword: auth_password,
			},
		},
	})

	return nil
}
