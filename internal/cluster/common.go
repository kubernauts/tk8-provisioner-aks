package cluster

import (
	"log"

	"github.com/kubernauts/tk8/pkg/common"
	"github.com/spf13/viper"
)

type AKSConfig struct {
	ClusterName       string
	Location          string
	AgentCount        string
	SSHKeyPath        string
	NetworkPlugin     string
	DNSPrefix         string
	ResourceGroupName string
	OSDiskSizeGB      string
	KubernetesVersion string
	VMSize            string
}

func GetAKSConfig() AKSConfig {
	ReadViperConfigFile("config")
	return AKSConfig{
		ClusterName:       viper.GetString("aks.cluster_name"),
		Location:          viper.GetString("aks.location"),
		AgentCount:        viper.GetString("aks.agent_count"),
		SSHKeyPath:        viper.GetString("aks.ssh_public_key"),
		NetworkPlugin:     viper.GetString("aks.network_plugin"),
		DNSPrefix:         viper.GetString("aks.dns_prefix"),
		ResourceGroupName: viper.GetString("aks.resource_group_name"),
		OSDiskSizeGB:      viper.GetString("aks.os_disk_size_gb"),
		KubernetesVersion: viper.GetString("aks.kubernetes_version"),
		VMSize:            viper.GetString("aks.vm_size"),
	}
}

func SetClusterName() {
	if len(common.Name) < 1 {
		config := GetAKSConfig()
		common.Name = config.ClusterName
	}
}

// ReadViperConfigFile is define the config paths and read the configuration file.
func ReadViperConfigFile(configName string) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("/tk8")
	verr := viper.ReadInConfig() // Find and read the config file.
	if verr != nil {             // Handle errors reading the config file.
		log.Fatalln(verr)
	}
}
