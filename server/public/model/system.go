// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"math/big"
)

const (
	SystemTelemetryId                      = "DiagnosticId"
	SystemRanUnitTests                     = "RanUnitTests"
	SystemLastSecurityTime                 = "LastSecurityTime"
	SystemActiveLicenseId                  = "ActiveLicenseId"
	SystemLastComplianceTime               = "LastComplianceTime"
	SystemAsymmetricSigningKeyKey          = "AsymmetricSigningKey"
	SystemPostActionCookieSecretKey        = "PostActionCookieSecret"
	SystemInstallationDateKey              = "InstallationDate"
	SystemOrganizationName                 = "OrganizationName"
	SystemFirstAdminRole                   = "FirstAdminRole"
	SystemFirstServerRunTimestampKey       = "FirstServerRunTimestamp"
	SystemClusterEncryptionKey             = "ClusterEncryptionKey"
	SystemUpgradedFromTeId                 = "UpgradedFromTE"
	SystemWarnMetricNumberOfTeams5         = "warn_metric_number_of_teams_5"
	SystemWarnMetricNumberOfChannels50     = "warn_metric_number_of_channels_50"
	SystemWarnMetricMfa                    = "warn_metric_mfa"
	SystemWarnMetricEmailDomain            = "warn_metric_email_domain"
	SystemWarnMetricNumberOfActiveUsers100 = "warn_metric_number_of_active_users_100"
	SystemWarnMetricNumberOfActiveUsers200 = "warn_metric_number_of_active_users_200"
	SystemWarnMetricNumberOfActiveUsers300 = "warn_metric_number_of_active_users_300"
	SystemWarnMetricNumberOfActiveUsers500 = "warn_metric_number_of_active_users_500"
	SystemWarnMetricNumberOfPosts2m        = "warn_metric_number_of_posts_2M"
	SystemWarnMetricLastRunTimestampKey    = "LastWarnMetricRunTimestamp"
	SystemFirstAdminVisitMarketplace       = "FirstAdminVisitMarketplace"
	SystemFirstAdminSetupComplete          = "FirstAdminSetupComplete"
	SystemLastAccessiblePostTime           = "LastAccessiblePostTime"
	SystemLastAccessibleFileTime           = "LastAccessibleFileTime"
	SystemHostedPurchaseNeedsScreening     = "HostedPurchaseNeedsScreening"
	AwsMeteringReportInterval              = 1
	AwsMeteringDimensionUsageHrs           = "UsageHrs"
	CloudRenewalEmail                      = "CloudRenewalEmail"
)

const (
	WarnMetricStatusLimitReached    = "true"
	WarnMetricStatusRunonce         = "runonce"
	WarnMetricStatusAck             = "ack"
	WarnMetricStatusStorePrefix     = "warn_metric_"
	WarnMetricJobInterval           = 24 * 7
	WarnMetricNumberOfActiveUsers25 = 25
	WarnMetricJobWaitTime           = 1000 * 3600 * 24 * 7 // 7 days
)

type System struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SystemPostActionCookieSecret struct {
	Secret []byte `json:"key,omitempty"`
}

type SystemAsymmetricSigningKey struct {
	ECDSAKey *SystemECDSAKey `json:"ecdsa_key,omitempty"`
}

type SystemECDSAKey struct {
	Curve string   `json:"curve"`
	X     *big.Int `json:"x"`
	Y     *big.Int `json:"y"`
	D     *big.Int `json:"d,omitempty"`
}

// ServerBusyState provides serialization for app.Busy.
type ServerBusyState struct {
	Busy      bool   `json:"busy"`
	Expires   int64  `json:"expires"`
	ExpiresTS string `json:"expires_ts,omitempty"`
}

type SupportPacket struct {
	GeneratedAt int64 `json:"generated_at"`

	/* Build information */

	ServerOS           string `yaml:"server_os"`
	ServerArchitecture string `yaml:"server_architecture"`
	ServerVersion      string `yaml:"server_version"`
	BuildHash          string `yaml:"build_hash,omitempty"`

	/* DB */

	DatabaseType          string `yaml:"database_type"`
	DatabaseVersion       string `yaml:"database_version"`
	DatabaseSchemaVersion string `yaml:"database_schema_version"`
	WebsocketConnections  int    `yaml:"websocket_connections"`
	MasterDbConnections   int    `yaml:"master_db_connections"`
	ReplicaDbConnections  int    `yaml:"read_db_connections"`

	/* Cluster */

	ClusterID string `yaml:"cluster_id"`

	/* File store */

	FileDriver string `yaml:"file_driver"`
	FileStatus string `yaml:"file_status"`

	/* LDAP */

	LdapVendorName    string `yaml:"ldap_vendor_name,omitempty"`
	LdapVendorVersion string `yaml:"ldap_vendor_version,omitempty"`

	/* Elastic Search */

	ElasticServerVersion string   `yaml:"elastic_server_version,omitempty"`
	ElasticServerPlugins []string `yaml:"elastic_server_plugins,omitempty"`

	/* License */

	LicenseTo             string `yaml:"license_to"`
	LicenseSupportedUsers int    `yaml:"license_supported_users,omitempty"`
	LicenseIsTrial        string `yaml:"license_is_trial,omitempty"`
	CustomerID            string `yaml:"customer_id"`
	LicenseID             string `yaml:"license_id"`
	ServerID              string `yaml:"server_id"`

	/* Server stats */

	ActiveUsers        int `yaml:"active_users"`
	DailyActiveUsers   int `yaml:"daily_active_users"`
	MonthlyActiveUsers int `yaml:"monthly_active_users"`
	InactiveUserCount  int `yaml:"inactive_user_count"`
	TotalPosts         int `yaml:"total_posts"`
	TotalChannels      int `yaml:"total_channels"`
	TotalTeams         int `yaml:"total_teams"`

	/* Jobs */

	DataRetentionJobs          []*Job `yaml:"data_retention_jobs"`
	MessageExportJobs          []*Job `yaml:"message_export_jobs"`
	ElasticPostIndexingJobs    []*Job `yaml:"elastic_post_indexing_jobs"`
	ElasticPostAggregationJobs []*Job `yaml:"elastic_post_aggregation_jobs"`
	BlevePostIndexingJobs      []*Job `yaml:"bleve_post_indexin_jobs"`
	LdapSyncJobs               []*Job `yaml:"ldap_sync_jobs"`
	MigrationJobs              []*Job `yaml:"migration_jobs"`
	ComplianceJobs             []*Job `yaml:"compliance_jobs"`
}

type FileData struct {
	Filename string
	Body     []byte
}

var WarnMetricsTable = map[string]WarnMetric{
	SystemWarnMetricMfa: {
		Id:        SystemWarnMetricMfa,
		Limit:     -1,
		IsBotOnly: true,
		IsRunOnce: true,
	},
	SystemWarnMetricEmailDomain: {
		Id:        SystemWarnMetricEmailDomain,
		Limit:     -1,
		IsBotOnly: true,
		IsRunOnce: true,
	},
	SystemWarnMetricNumberOfTeams5: {
		Id:        SystemWarnMetricNumberOfTeams5,
		Limit:     5,
		IsBotOnly: true,
		IsRunOnce: true,
	},
	SystemWarnMetricNumberOfChannels50: {
		Id:        SystemWarnMetricNumberOfChannels50,
		Limit:     50,
		IsBotOnly: true,
		IsRunOnce: true,
	},
	SystemWarnMetricNumberOfActiveUsers100: {
		Id:        SystemWarnMetricNumberOfActiveUsers100,
		Limit:     100,
		IsBotOnly: true,
		IsRunOnce: true,
	},
	SystemWarnMetricNumberOfActiveUsers200: {
		Id:        SystemWarnMetricNumberOfActiveUsers200,
		Limit:     200,
		IsBotOnly: true,
		IsRunOnce: true,
	},
	SystemWarnMetricNumberOfActiveUsers300: {
		Id:        SystemWarnMetricNumberOfActiveUsers300,
		Limit:     300,
		IsBotOnly: true,
		IsRunOnce: true,
	},
	SystemWarnMetricNumberOfActiveUsers500: {
		Id:        SystemWarnMetricNumberOfActiveUsers500,
		Limit:     500,
		IsBotOnly: false,
		IsRunOnce: true,
	},
	SystemWarnMetricNumberOfPosts2m: {
		Id:        SystemWarnMetricNumberOfPosts2m,
		Limit:     2000000,
		IsBotOnly: false,
		IsRunOnce: true,
	},
}

type WarnMetric struct {
	Id         string
	Limit      int64
	IsBotOnly  bool
	IsRunOnce  bool
	SkipAction bool
}

type WarnMetricDisplayTexts struct {
	BotTitle          string
	BotMessageBody    string
	BotSuccessMessage string
	EmailBody         string
}
type WarnMetricStatus struct {
	Id          string `json:"id"`
	Limit       int64  `json:"limit"`
	Acked       bool   `json:"acked"`
	StoreStatus string `json:"store_status,omitempty"`
}

type SendWarnMetricAck struct {
	ForceAck bool `json:"forceAck"`
}

type AppliedMigration struct {
	Version int    `json:"version"`
	Name    string `json:"name"`
}

type LogFilter struct {
	ServerNames []string `json:"server_names"`
	LogLevels   []string `json:"log_levels"`
	DateFrom    string   `json:"date_from"`
	DateTo      string   `json:"date_to"`
}

type LogEntry struct {
	Timestamp string
	Level     string
}
