package cmds

import (
	"context"
	"time"

	"github.com/appscode/go/flags"
	"github.com/appscode/go/log"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
	"github.com/appscode/pharmer/config"
	"github.com/appscode/pharmer/phid"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCmdCreate() *cobra.Command {
	cluster := &api.Cluster{}
	nodes := map[string]int{}

	cmd := &cobra.Command{
		Use:               "create",
		Short:             "Create a Kubernetes cluster for a given cloud provider",
		Example:           "create --provider=(aws|gce|cc) --nodes=t1=1,t2=2 --zone=us-central1-f demo-cluster",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			flags.EnsureRequiredFlags(cmd, "provider", "zone", "nodes")

			if len(args) == 0 {
				log.Fatalln("Missing cluster name")
			}
			if len(args) > 1 {
				log.Fatalln("Multiple cluster name provided.")
			}
			cluster.Name = args[0]
			cfgFile, _ := config.GetConfigFile(cmd.Flags())
			cfg, err := config.LoadConfig(cfgFile)
			if err != nil {
				log.Fatalln(err)
			}
			ctx := cloud.NewContext(context.TODO(), cfg)
			cluster, err = cloud.Create(ctx, cluster)
			if err != nil {
				log.Fatalln(err)
			}

			for sku, count := range nodes {
				ig := api.InstanceGroup{
					ObjectMeta: metav1.ObjectMeta{
						Name:              sku + "-pool",
						UID:               phid.NewInstanceGroup(),
						CreationTimestamp: metav1.Time{Time: time.Now()},
						Labels: map[string]string{
							"node-role.kubernetes.io/node": "true",
						},
					},
					Spec: api.InstanceGroupSpec{
						SKU:           sku,
						Count:         int64(count),
						SpotInstances: false,
						//DiskType:      "",
						//DiskSize:      0,
					},
				}
				_, err := cloud.Store(ctx).InstanceGroups(cluster.Name).Create(&ig)
				if err != nil {
					log.Fatalln(err)
				}
			}
		},
	}

	cmd.Flags().StringVar(&cluster.Spec.Cloud.CloudProvider, "provider", "", "Provider name")
	cmd.Flags().StringVar(&cluster.Spec.Cloud.Zone, "zone", "", "Cloud provider zone name")
	cmd.Flags().StringVar(&cluster.Spec.CredentialName, "credential-uid", "", "Use preconfigured cloud credential uid")
	cmd.Flags().StringVar(&cluster.Spec.KubernetesVersion, "kubernetes-version", "", "Kubernetes version")
	cmd.Flags().StringVar(&cluster.Spec.KubeadmVersion, "kubeadm-version", "", "Kubeadm version")
	cmd.Flags().BoolVar(&cluster.Spec.DoNotDelete, "do-not-delete", false, "Set do not delete flag")
	cmd.Flags().StringToIntVar(&nodes, "nodes", map[string]int{}, "Node set configuration")

	return cmd
}