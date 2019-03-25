package cluster

import (
	"log"
	"os"
	"os/exec"

	"github.com/kubernauts/tk8/pkg/common"
	"github.com/kubernauts/tk8/pkg/provisioner"
	"github.com/kubernauts/tk8/pkg/templates"
)

func aksPrepareConfigFiles(Name string) {
	templates.ParseTemplate(templates.VariablesAKS, "./inventory/"+common.Name+"/provisioner/variables.tf", GetAKSConfig())
	templates.ParseTemplate(templates.Credentials, "./inventory/"+common.Name+"/provisioner/credentials.tfvars", common.GetCredentials())

}

// Install is used to setup the Kubernetes Cluster (AKS)
func Install() {
	os.MkdirAll("./inventory/"+common.Name+"/provisioner/modules/aks", 0755)
	exec.Command("cp", "-rfp", "./provisioner/aks/", "./inventory/"+common.Name+"/provisioner").Run()
	aksPrepareConfigFiles(common.Name)

	// Make shell scripts executable
	aksShellScriptsPath := "./inventory/" + common.Name + "/provisioner/modules/aks"
	if err := os.Chmod(aksShellScriptsPath+"/clean-network.sh", 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.Chmod(aksShellScriptsPath+"/config-network.sh", 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.Chmod(aksShellScriptsPath+"/helm-install.sh", 0755); err != nil {
		log.Fatal(err)
	}
	// Check if a terraform state file already exists
	if _, err := os.Stat("./inventory/" + common.Name + "/provisioner/terraform.tfstate"); err == nil {
		log.Fatal("There is an existing cluster, please remove terraform.tfstate file or delete the installation before proceeding")
	} else {
		log.Println("starting terraform init")

		provisioner.ExecuteTerraform("init", "./inventory/"+common.Name+"/provisioner/")

	}

	provisioner.ExecuteTerraform("apply", "./inventory/"+common.Name+"/provisioner/")

	log.Println("Kubeconfig file can be found in the output results starting with kube_config variable")
	log.Println("Voila! AKS cluster is up and running")

	os.Exit(0)

}

func Upgrade() {

	// Check if a terraform state file already exists
	if _, err := os.Stat("./inventory/" + common.Name + "/provisioner/terraform.tfstate"); err == nil {
		if os.IsNotExist(err) {
			log.Fatal("No terraform.tfstate file found. Upgrade can only be done on an existing cluster.")
		}
	}
	log.Println("Starting Upgrade of the existing cluster")
	provisioner.ExecuteTerraform("apply", "./inventory/"+common.Name+"/provisioner/")

	log.Println("Kubeconfig file can be found in the output results starting with kube_config variable")
	log.Println("Voila! AKS cluster is up and running")

	os.Exit(0)

}

// Reset is used to reset the  Kubernetes Cluster back to rollout on the infrastructure.
func AKSReset() {
	provisioner.NotImplemented()
}

func AKSDestroy() {
	log.Println("Starting deletion of AKS cluster")
	log.Println("starting terraform destroy")
	provisioner.ExecuteTerraform("destroy", "./inventory/"+common.Name+"/provisioner/")
}
