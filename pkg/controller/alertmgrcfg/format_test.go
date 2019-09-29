package alertmgrcfg

/*
 Copyright [2019] [Platform9 Systems, Inc]

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"reflect"
	"testing"

	monitoringv1alpha1 "github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestGetFormatter(t *testing.T) {
	var f format
	f, err := getFormatter("slack")
	assert.Equal(t, nil, err)
	assert.Equal(t, reflect.TypeOf(f), reflect.TypeOf(slackconfig{}))

	f, err = getFormatter("email")
	assert.Equal(t, nil, err)
	assert.Equal(t, reflect.TypeOf(f), reflect.TypeOf(emailconfig{}))
}

func TestSlackFormat(t *testing.T) {
	URL := "https://hooks.slack.com/services/xxx/yyy"
	channel := "#alertmgr"
	severity := "info"

	amc := monitoringv1alpha1.AlertMgrCfg{
		Spec: monitoringv1alpha1.AlertMgrCfgSpec{
			Type: "slack",
			Params: []monitoringv1alpha1.Param{
				monitoringv1alpha1.Param{
					Name:  "url",
					Value: URL,
				},
				monitoringv1alpha1.Param{
					Name:  "channel",
					Value: channel,
				},
				monitoringv1alpha1.Param{
					Name:  "severity",
					Value: severity,
				},
			},
		},
	}

	acfg := alertConfig{
		Receivers: []receiver{
			receiver{
				Name: "webhook",
			},
		},
	}
	err := formatReceiver(&amc, &acfg)
	assert.Equal(t, nil, err)

	assert.Equal(t, acfg.Receivers[1].SlackConfigs[0].ApiURL, URL)
	assert.Equal(t, acfg.Receivers[1].SlackConfigs[0].Channel, channel)
	assert.Equal(t, acfg.Route.Routes[0].MatchRe["severity"], severity)
}

func TestEmailFormat(t *testing.T) {
	to := "to@p9.com"
	from := "from@p9.com"
	smarthost := "p9.local:8887"
	auth_username := "from@p9.com"
	auth_identity := "from@p9.com"
	auth_password := "pwd"
	severity := "info"

	amc := monitoringv1alpha1.AlertMgrCfg{
		Spec: monitoringv1alpha1.AlertMgrCfgSpec{
			Type: "email",
			Params: []monitoringv1alpha1.Param{
				monitoringv1alpha1.Param{
					Name:  "to",
					Value: to,
				},
				monitoringv1alpha1.Param{
					Name:  "from",
					Value: from,
				},
				monitoringv1alpha1.Param{
					Name:  "smarthost",
					Value: smarthost,
				},
				monitoringv1alpha1.Param{
					Name:  "auth_username",
					Value: auth_username,
				},
				monitoringv1alpha1.Param{
					Name:  "auth_identity",
					Value: auth_identity,
				},
				monitoringv1alpha1.Param{
					Name:  "auth_password",
					Value: auth_password,
				},
				monitoringv1alpha1.Param{
					Name:  "severity",
					Value: severity,
				},
			},
		},
	}

	acfg := alertConfig{
		Receivers: []receiver{
			receiver{
				Name: "webhook",
			},
		},
	}

	err := formatReceiver(&amc, &acfg)
	assert.Equal(t, nil, err)

	assert.Equal(t, acfg.Receivers[1].EmailConfigs[0].To, to)
	assert.Equal(t, acfg.Receivers[1].EmailConfigs[0].From, from)
	assert.Equal(t, acfg.Receivers[1].EmailConfigs[0].SmartHost, smarthost)
	assert.Equal(t, acfg.Receivers[1].EmailConfigs[0].AuthUsername, auth_username)
	assert.Equal(t, acfg.Receivers[1].EmailConfigs[0].AuthIdentity, auth_identity)
	assert.Equal(t, acfg.Receivers[1].EmailConfigs[0].AuthPassword, auth_password)
	assert.Equal(t, acfg.Route.Routes[0].MatchRe["severity"], severity)
}
