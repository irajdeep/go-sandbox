package main

import (
	"fmt"

	"k8s.io/api/admissionregistration/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	url := "https://127.0.0.1"

	config := &v1beta1.ValidatingWebhookConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-validation-webhook-configration",
		},
		Webhooks: []v1beta1.ValidatingWebhook{
			{
				Name: "pod-validation.booking-admission-controller.booking.com",
				ClientConfig: v1beta1.WebhookClientConfig{
					URL: &url,
				},
				Rules: []v1beta1.RuleWithOperations{{
					Operations: []v1beta1.OperationType{
						v1beta1.Create,
						v1beta1.Update,
					},
					Rule: v1beta1.Rule{
						APIVersions: []string{"*"},
						APIGroups:   []string{"*"},
						Resources:   []string{"pods"},
					},
				},
				},
			},
		},
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s", string(data))

}
