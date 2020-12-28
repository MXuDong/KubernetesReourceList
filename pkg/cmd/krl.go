package cmd

import (
	"fmt"
	"github.com/A-Donga/TablePrinter"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
)

type KRLOptions struct {
	configFlags *genericclioptions.ConfigFlags

	namespace string
	args      []string

	genericclioptions.IOStreams
}

func NewKRLOptions(streams genericclioptions.IOStreams) *KRLOptions {
	return &KRLOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
	}
}

func NewCmdKRL(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewKRLOptions(streams)

	cmd := &cobra.Command{
		Use:          "rl [resource] [other args[namespace]...]",
		Short:        "Show some resource in table",
		Example:      "kubectl rl all",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}
			return nil
		},
	}

	o.configFlags.AddFlags(cmd.Flags())
	return cmd
}

func (o *KRLOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args

	return nil
}

func (o *KRLOptions) Validate() error {

	return nil
}

func (o *KRLOptions) Run() error {
	builder := resource.NewBuilder(o.configFlags)
	result := builder.Unstructured().
		NamespaceParam(*o.configFlags.Namespace).DefaultNamespace().AllNamespaces(false).
		ResourceTypeOrNameArgs(false, o.args...).
		SelectAllParam(true).
		Flatten().
		Do()

	table := TablePrinter.InitTitle([]string{"name", "namespace", "api-group", "api-kind"}, nil)
	err := result.Visit(func(info *resource.Info, err error) error {
		name := info.Name
		namespace := info.Namespace
		apiGroup := info.Mapping.GroupVersionKind.Group
		apiKind := info.Mapping.GroupVersionKind.Kind
		table.AddRow([] string {name, namespace, apiGroup, apiKind})
		return nil
	})
	fmt.Println(table.GetTable())

	return err
}