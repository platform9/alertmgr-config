package alertmgrcfg

import (
	"os"

	monitoringv1alpha1 "github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1"
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
	var url, channel string
	for _, param := range amc.Spec.Params {
		switch param.Name {
		case "url":
			url = param.Value
		case "channel":
			channel = param.Value
		}
	}

	if url == "" {
		return os.ErrInvalid
	}

	if channel == "" {
		return os.ErrInvalid
	}

	acfg.Receivers[0].SlackConfigs = append(acfg.Receivers[0].SlackConfigs,
		slackconfig{
			ApiURL:  url,
			Channel: channel,
		})

	return nil
}

func (f emailconfig) formatAlert(amc *monitoringv1alpha1.AlertMgrCfg, acfg *alertConfig) error {
	var to, from, smarthost string
	for _, param := range amc.Spec.Params {
		switch param.Name {
		case "to":
			to = param.Value
		case "from":
			from = param.Value
		case "smarthost":
			smarthost = param.Value
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

	acfg.Receivers[0].EmailConfigs = append(acfg.Receivers[0].EmailConfigs,
		emailconfig{
			To:        to,
			From:      from,
			SmartHost: smarthost,
		})

	return nil
}
