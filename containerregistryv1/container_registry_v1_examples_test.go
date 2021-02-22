// +build examples

/**
 * (C) Copyright IBM Corp. 2020, 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package containerregistryv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
	"github.com/IBM/go-sdk-core/v4/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const externalConfigFile = "../container_registry_v1.env"

var (
	containerRegistryService *containerregistryv1.ContainerRegistryV1
	config                   map[string]string
	configLoaded             bool = false
)

// Globlal variables to hold link values
var (
	namespaceLink string
)

func shouldSkipTest() {
	Skip("Container Registry examples are not intended to be runnable tests")
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`ContainerRegistryV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}
			/**
			Your configuration file (container_registry_v1.env) should contain the following variables.
			CONTAINER_REGISTRY_URL=[Registry URL, eg https://uk.icr.io]
			CONTAINER_REGISTRY_AUTH_TYPE=iam
			CONTAINER_REGISTRY_AUTH_URL=https://iam.cloud.ibm.com/identity/token
			CONTAINER_REGISTRY_APIKEY=[An IAM Apikey]
			*/
			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(containerregistryv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			containerRegistryServiceOptions := &containerregistryv1.ContainerRegistryV1Options{
				Account: core.StringPtr("accountID"),
			}

			containerRegistryService, err = containerregistryv1.NewContainerRegistryV1UsingExternalConfig(containerRegistryServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(containerRegistryService).ToNot(BeNil())
		})
	})

	Describe(`ContainerRegistryV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNamespace request example`, func() {
			// begin-create_namespace

			createNamespaceOptions := containerRegistryService.NewCreateNamespaceOptions(
				"my_example_namespace",
			)

			namespace, response, err := containerRegistryService.CreateNamespace(createNamespaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(namespace, "", "  ")
			fmt.Println(string(b))

			// end-create_namespace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Or(Equal(201), Equal(200)))
			Expect(namespace).ToNot(BeNil())

			namespaceLink = *namespace.Namespace

		})
		It(`GetAuth request example`, func() {
			// begin-get_auth

			getAuthOptions := containerRegistryService.NewGetAuthOptions()

			authOptions, response, err := containerRegistryService.GetAuth(getAuthOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(authOptions, "", "  ")
			fmt.Println(string(b))

			// end-get_auth

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(authOptions).ToNot(BeNil())

		})
		It(`UpdateAuth request example`, func() {
			// begin-update_auth

			updateAuthOptions := containerRegistryService.NewUpdateAuthOptions()
			updateAuthOptions.SetIamAuthz(true)

			response, err := containerRegistryService.UpdateAuth(updateAuthOptions)
			if err != nil {
				panic(err)
			}

			// end-update_auth

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`GetSettings request example`, func() {
			// begin-get_settings

			getSettingsOptions := containerRegistryService.NewGetSettingsOptions()

			accountSettings, response, err := containerRegistryService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
		It(`UpdateSettings request example`, func() {
			// begin-update_settings

			updateSettingsOptions := containerRegistryService.NewUpdateSettingsOptions()
			updateSettingsOptions.SetPlatformMetrics(true)

			response, err := containerRegistryService.UpdateSettings(updateSettingsOptions)
			if err != nil {
				panic(err)
			}

			// end-update_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListImages request example`, func() {
			// begin-list_images

			listImagesOptions := containerRegistryService.NewListImagesOptions()

			remoteApiImage, response, err := containerRegistryService.ListImages(listImagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(remoteApiImage, "", "  ")
			fmt.Println(string(b))

			// end-list_images

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(remoteApiImage).ToNot(BeNil())

		})
		It(`BulkDeleteImages request example`, func() {
			// begin-bulk_delete_images

			bulkDeleteImagesOptions := containerRegistryService.NewBulkDeleteImagesOptions(
				[]string{"testString"},
			)

			imageBulkDeleteResult, response, err := containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageBulkDeleteResult, "", "  ")
			fmt.Println(string(b))

			// end-bulk_delete_images

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageBulkDeleteResult).ToNot(BeNil())

		})
		It(`ListImageDigests request example`, func() {
			// begin-list_image_digests

			listImageDigestsOptions := containerRegistryService.NewListImageDigestsOptions()
			listImageDigestsOptions.SetExcludeTagged(false)
			listImageDigestsOptions.SetExcludeVa(false)
			listImageDigestsOptions.SetIncludeIBM(false)

			digestListImage, response, err := containerRegistryService.ListImageDigests(listImageDigestsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(digestListImage, "", "  ")
			fmt.Println(string(b))

			// end-list_image_digests

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(digestListImage).ToNot(BeNil())

		})
		It(`TagImage request example`, func() {
			// begin-tag_image

			tagImageOptions := containerRegistryService.NewTagImageOptions(
				"testString",
				"testString",
			)

			response, err := containerRegistryService.TagImage(tagImageOptions)
			if err != nil {
				panic(err)
			}

			// end-tag_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
		It(`InspectImage request example`, func() {
			// begin-inspect_image

			inspectImageOptions := containerRegistryService.NewInspectImageOptions(
				"testString",
			)

			imageInspection, response, err := containerRegistryService.InspectImage(inspectImageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageInspection, "", "  ")
			fmt.Println(string(b))

			// end-inspect_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageInspection).ToNot(BeNil())

		})
		It(`GetImageManifest request example`, func() {
			// begin-get_image_manifest

			getImageManifestOptions := containerRegistryService.NewGetImageManifestOptions(
				"testString",
			)

			response, err := containerRegistryService.GetImageManifest(getImageManifestOptions)
			if err != nil {
				panic(err)
			}

			// end-get_image_manifest

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`GetMessages request example`, func() {
			// begin-get_messages

			getMessagesOptions := containerRegistryService.NewGetMessagesOptions()

			getMessagesResponse, response, err := containerRegistryService.GetMessages(getMessagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getMessagesResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_messages

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getMessagesResponse).ToNot(BeNil())

		})
		It(`ListNamespaces request example`, func() {
			// begin-list_namespaces

			listNamespacesOptions := containerRegistryService.NewListNamespacesOptions()

			result, response, err := containerRegistryService.ListNamespaces(listNamespacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-list_namespaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`ListNamespaceDetails request example`, func() {
			// begin-list_namespace_details

			listNamespaceDetailsOptions := containerRegistryService.NewListNamespaceDetailsOptions()

			namespaceDetail, response, err := containerRegistryService.ListNamespaceDetails(listNamespaceDetailsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(namespaceDetail, "", "  ")
			fmt.Println(string(b))

			// end-list_namespace_details

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespaceDetail).ToNot(BeNil())

		})
		It(`AssignNamespace request example`, func() {
			// begin-assign_namespace

			assignNamespaceOptions := containerRegistryService.NewAssignNamespaceOptions(
				"testString",
				"testString",
			)

			namespace, response, err := containerRegistryService.AssignNamespace(assignNamespaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(namespace, "", "  ")
			fmt.Println(string(b))

			// end-assign_namespace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespace).ToNot(BeNil())

		})
		It(`GetPlans request example`, func() {
			// begin-get_plans

			getPlansOptions := containerRegistryService.NewGetPlansOptions()

			plan, response, err := containerRegistryService.GetPlans(getPlansOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(plan, "", "  ")
			fmt.Println(string(b))

			// end-get_plans

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(plan).ToNot(BeNil())

		})
		It(`UpdatePlans request example`, func() {
			// begin-update_plans

			updatePlansOptions := containerRegistryService.NewUpdatePlansOptions()
			updatePlansOptions.SetPlan("Standard")

			response, err := containerRegistryService.UpdatePlans(updatePlansOptions)
			if err != nil {
				panic(err)
			}

			// end-update_plans

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`GetQuota request example`, func() {
			// begin-get_quota

			getQuotaOptions := containerRegistryService.NewGetQuotaOptions()

			quota, response, err := containerRegistryService.GetQuota(getQuotaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(quota, "", "  ")
			fmt.Println(string(b))

			// end-get_quota

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(quota).ToNot(BeNil())

		})
		It(`UpdateQuota request example`, func() {
			// begin-update_quota

			updateQuotaOptions := containerRegistryService.NewUpdateQuotaOptions()
			updateQuotaOptions.SetTrafficMegabytes(int64(480))

			response, err := containerRegistryService.UpdateQuota(updateQuotaOptions)
			if err != nil {
				panic(err)
			}

			// end-update_quota

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`ListRetentionPolicies request example`, func() {
			// begin-list_retention_policies

			listRetentionPoliciesOptions := containerRegistryService.NewListRetentionPoliciesOptions()

			mapStringRetentionPolicy, response, err := containerRegistryService.ListRetentionPolicies(listRetentionPoliciesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(mapStringRetentionPolicy, "", "  ")
			fmt.Println(string(b))

			// end-list_retention_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringRetentionPolicy).ToNot(BeNil())

		})
		It(`SetRetentionPolicy request example`, func() {
			// begin-set_retention_policy

			setRetentionPolicyOptions := containerRegistryService.NewSetRetentionPolicyOptions("birds")
			setRetentionPolicyOptions.SetImagesPerRepo(int64(10))
			setRetentionPolicyOptions.SetRetainUntagged(false)

			response, err := containerRegistryService.SetRetentionPolicy(setRetentionPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-set_retention_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`AnalyzeRetentionPolicy request example`, func() {
			// begin-analyze_retention_policy

			analyzeRetentionPolicyOptions := containerRegistryService.NewAnalyzeRetentionPolicyOptions("birds")
			analyzeRetentionPolicyOptions.SetImagesPerRepo(int64(10))
			analyzeRetentionPolicyOptions.SetRetainUntagged(false)

			mapStringstring, response, err := containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(mapStringstring, "", "  ")
			fmt.Println(string(b))

			// end-analyze_retention_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringstring).ToNot(BeNil())

		})
		It(`GetRetentionPolicy request example`, func() {
			// begin-get_retention_policy

			getRetentionPolicyOptions := containerRegistryService.NewGetRetentionPolicyOptions(
				"testString",
			)

			retentionPolicy, response, err := containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(retentionPolicy, "", "  ")
			fmt.Println(string(b))

			// end-get_retention_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(retentionPolicy).ToNot(BeNil())

		})
		It(`ListDeletedImages request example`, func() {
			// begin-list_deleted_images

			listDeletedImagesOptions := containerRegistryService.NewListDeletedImagesOptions()

			mapStringTrash, response, err := containerRegistryService.ListDeletedImages(listDeletedImagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(mapStringTrash, "", "  ")
			fmt.Println(string(b))

			// end-list_deleted_images

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringTrash).ToNot(BeNil())

		})
		It(`RestoreTags request example`, func() {
			// begin-restore_tags

			restoreTagsOptions := containerRegistryService.NewRestoreTagsOptions(
				"testString",
			)

			restoreResult, response, err := containerRegistryService.RestoreTags(restoreTagsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(restoreResult, "", "  ")
			fmt.Println(string(b))

			// end-restore_tags

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(restoreResult).ToNot(BeNil())

		})
		It(`RestoreImage request example`, func() {
			// begin-restore_image

			restoreImageOptions := containerRegistryService.NewRestoreImageOptions(
				"testString",
			)

			response, err := containerRegistryService.RestoreImage(restoreImageOptions)
			if err != nil {
				panic(err)
			}

			// end-restore_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteNamespace request example`, func() {
			// begin-delete_namespace

			deleteNamespaceOptions := containerRegistryService.NewDeleteNamespaceOptions(
				namespaceLink,
			)

			response, err := containerRegistryService.DeleteNamespace(deleteNamespaceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_namespace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteImageTag request example`, func() {
			// begin-delete_image_tag

			deleteImageTagOptions := containerRegistryService.NewDeleteImageTagOptions(
				"testString",
			)

			imageDeleteResult, response, err := containerRegistryService.DeleteImageTag(deleteImageTagOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageDeleteResult, "", "  ")
			fmt.Println(string(b))

			// end-delete_image_tag

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDeleteResult).ToNot(BeNil())

		})
		It(`DeleteImage request example`, func() {
			// begin-delete_image

			deleteImageOptions := containerRegistryService.NewDeleteImageOptions(
				"testString",
			)

			imageDeleteResult, response, err := containerRegistryService.DeleteImage(deleteImageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageDeleteResult, "", "  ")
			fmt.Println(string(b))

			// end-delete_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDeleteResult).ToNot(BeNil())

		})
	})
})
