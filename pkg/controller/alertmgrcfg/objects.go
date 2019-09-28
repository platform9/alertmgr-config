package alertmgrcfg

type alertConfig struct {
	Global    global     `yaml:"global"`
	Route     route      `yaml:"route"`
	Receivers []receiver `yaml:"receivers"`
}

type global struct {
	ResolveTimeout string `yaml:"resolve_timeout"`
}

type route struct {
	GroupBy        []string `yaml:"group_by"`
	GroupWait      string   `yaml:"group_wait"`
	GroupInterval  string   `yaml:"group_interval"`
	RepeatInterval string   `yaml:"repeat_interval"`
	Receiver       string   `yaml:"receiver"`
}

type slackconfig struct {
	ApiURL  string `yaml:"api_url"`
	Channel string `yaml:"channel"`
}

type emailconfig struct {
	To        string `yaml:"to"`
	From      string `yaml:"from"`
	SmartHost string `yaml:"smarthost"`
}

type receiver struct {
	Name         string        `yaml:"name"`
	SlackConfigs []slackconfig `yaml:"slack_configs,omitempty"`
	EmailConfigs []emailconfig `yaml:"email_configs,omitempty"`
}
