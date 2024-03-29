//go:build integration
// +build integration

/**
 * (C) Copyright IBM Corp. 2020, 2023.
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
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the containerregistryv1 package.
 *
 * Notes:
 *
 * Your configuration file (container_registry_v1.env) should contain the following variables.
CONTAINER_REGISTRY_URL=[Registry URL, eg https://uk.icr.io]
CONTAINER_REGISTRY_AUTH_TYPE=iam
CONTAINER_REGISTRY_AUTH_URL=https://iam.cloud.ibm.com/identity/token
CONTAINER_REGISTRY_APIKEY=[An IAM Apikey]
CONTAINER_REGISTRY_ACCOUNT_ID=[Your test account ID]
CONTAINER_REGISTRY_RESOURCE_GROUP_ID=[Your resource group ID]
CONTAINER_REGISTRY_NAMESPACE=[Namespace name, to be created and deleted by the test]
CONTAINER_REGISTRY_SEED_IMAGE=[An existing namespace/repo:tag to copy in this test, eg: my_existing_namespace/seedimage:1234]
CONTAINER_REGISTRY_SEED_DIGEST=[The digest of the seed image, eg: sha256:aaaaaa9e4044327fd101ca1fd4043e6f3ad921ae7ee901e9142e6e36deadbeef]
 *
 * The integration test will automatically skip tests if the required config file is not available.
*/

var _ = Describe(`ContainerRegistryV1 Integration Tests`, func() {
	const externalConfigFile = "../container_registry_v1.env"

	var (
		err                      error
		containerRegistryService *containerregistryv1.ContainerRegistryV1
		serviceURL               string
		baseNamespace            string
		accountID                string
		resouceGroupID           string
		registryDNSName          string
		seedImage                string
		seedDigest               string
		config                   map[string]string

		// Variables to hold link values
		namespaceLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(containerregistryv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}
			baseNamespace = config["NAMESPACE"]
			if baseNamespace == "" {
				Skip("Unable to load baseNamespace configuration property, skipping tests")
			}
			accountID = config["ACCOUNT_ID"]
			if accountID == "" {
				Skip("Unable to load accountID configuration property, skipping tests")
			}
			resouceGroupID = config["RESOURCE_GROUP_ID"]
			if resouceGroupID == "" {
				Skip("Unable to load resouceGroupID configuration property, skipping tests")
			}
			registryDNSName = strings.TrimPrefix(serviceURL, "https://")
			seedImage = config["SEED_IMAGE"]
			if seedImage == "" {
				Skip("Unable to load seedImage configuration property, skipping tests")
			}
			seedDigest = config["SEED_DIGEST"]
			if seedDigest == "" {
				Skip("Unable to load seedDigest configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			containerRegistryServiceOptions := &containerregistryv1.ContainerRegistryV1Options{
				Account: core.StringPtr(accountID),
			}

			containerRegistryService, err = containerregistryv1.NewContainerRegistryV1UsingExternalConfig(containerRegistryServiceOptions)
			Expect(err).To(BeNil())
			Expect(containerRegistryService).ToNot(BeNil())
			Expect(containerRegistryService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			containerRegistryService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateNamespace - Create namespace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNamespace(createNamespaceOptions *CreateNamespaceOptions)`, func() {
			createNamespaceOptions := &containerregistryv1.CreateNamespaceOptions{
				Name: core.StringPtr(baseNamespace),
			}

			namespace, response, err := containerRegistryService.CreateNamespace(createNamespaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Or(Equal(201), Equal(200)))
			Expect(*namespace.Namespace).To(Equal(baseNamespace))

			namespaceLink = *namespace.Namespace
			fmt.Fprintf(GinkgoWriter, "Saved namespaceLink value: %v\n", namespaceLink)
		})
	})

	Describe(`GetAuth - Get authorization options`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAuth(getAuthOptions *GetAuthOptions)`, func() {
			getAuthOptions := &containerregistryv1.GetAuthOptions{}

			authOptions, response, err := containerRegistryService.GetAuth(getAuthOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(*authOptions.IamAuthz).To(BeTrue())
		})
	})

	Describe(`UpdateAuth - Update authorization options`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAuth(updateAuthOptions *UpdateAuthOptions)`, func() {
			updateAuthOptions := &containerregistryv1.UpdateAuthOptions{
				IamAuthz:    core.BoolPtr(true),
				PrivateOnly: core.BoolPtr(false),
			}

			response, err := containerRegistryService.UpdateAuth(updateAuthOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`GetSettings - Get account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {

			getSettingsOptions := &containerregistryv1.GetSettingsOptions{}

			accountSettings, response, err := containerRegistryService.GetSettings(getSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
	})

	Describe(`UpdateSettings - Update account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {

			updateSettingsOptions := &containerregistryv1.UpdateSettingsOptions{
				PlatformMetrics: core.BoolPtr(false),
			}

			response, err := containerRegistryService.UpdateSettings(updateSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`TagImage - Create tag`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TagImage(tagImageOptions *TagImageOptions)`, func() {

			tagImageOptions := &containerregistryv1.TagImageOptions{
				Fromimage: core.StringPtr(fmt.Sprintf("%s/%s", registryDNSName, seedImage)),
				Toimage:   core.StringPtr(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)),
			}

			response, err := containerRegistryService.TagImage(tagImageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
	})

	Describe(`ListImages - List images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListImages(listImagesOptions *ListImagesOptions)`, func() {
			listImagesOptions := &containerregistryv1.ListImagesOptions{
				Namespace:            core.StringPtr(namespaceLink),
				IncludeIBM:           core.BoolPtr(false),
				IncludePrivate:       core.BoolPtr(true),
				IncludeManifestLists: core.BoolPtr(true),
				Vulnerabilities:      core.BoolPtr(true),
			}

			remoteAPIImage, response, err := containerRegistryService.ListImages(listImagesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(remoteAPIImage).ToNot(BeNil())
			Expect(remoteAPIImage[0].RepoTags[0]).To(Equal(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)))
		})
	})

	Describe(`BulkDeleteImages - Bulk delete images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`BulkDeleteImages(bulkDeleteImagesOptions *BulkDeleteImagesOptions)`, func() {
			bulkDeleteImagesOptions := &containerregistryv1.BulkDeleteImagesOptions{
				BulkDelete: []string{fmt.Sprintf("%s/%s/notexist:1", registryDNSName, namespaceLink)},
			}

			imageBulkDeleteResult, response, err := containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageBulkDeleteResult).ToNot(BeNil())
		})
	})

	Describe(`ListImageDigests - List images by digest`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListImageDigests(listImageDigestsOptions *ListImageDigestsOptions)`, func() {
			listImageDigestsOptions := &containerregistryv1.ListImageDigestsOptions{
				ExcludeTagged: core.BoolPtr(false),
				ExcludeVa:     core.BoolPtr(false),
				IncludeIBM:    core.BoolPtr(false),
			}

			imageDigest, response, err := containerRegistryService.ListImageDigests(listImageDigestsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDigest).ToNot(BeNil())
			found := false
			for _, img := range imageDigest {
				if img.RepoTags[fmt.Sprintf("%s/%s/sdktest", registryDNSName, namespaceLink)] != nil {
					found = true
				}
			}
			Expect(found).To(BeTrue())
		})
	})

	Describe(`InspectImage - Inspect an image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InspectImage(inspectImageOptions *InspectImageOptions)`, func() {
			inspectImageOptions := &containerregistryv1.InspectImageOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)),
			}

			imageInspection, response, err := containerRegistryService.InspectImage(inspectImageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageInspection).ToNot(BeNil())
		})
	})

	Describe(`GetImageManifest - Get image manifest`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetImageManifest(getImageManifestOptions *GetImageManifestOptions)`, func() {
			getImageManifestOptions := &containerregistryv1.GetImageManifestOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)),
			}

			// The result variable is a map[string]interface{} containing the unmarshalled manifest
			result, response, err := containerRegistryService.GetImageManifest(getImageManifestOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			Expect(result["schemaVersion"]).To(Equal(float64(2)))

			contentType := response.Headers.Get("Content-Type")
			Expect(contentType).To(Equal("application/vnd.docker.distribution.manifest.v2+json"))
		})
	})

	Describe(`GetMessages - Get messages`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMessages(getMessagesOptions *GetMessagesOptions)`, func() {
			getMessagesOptions := &containerregistryv1.GetMessagesOptions{}

			result, response, err := containerRegistryService.GetMessages(getMessagesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Or(Equal(200), Equal(204)))
			if result != nil {
				Expect(*result).ToNot(BeEmpty())
			}

		})
	})

	Describe(`ListNamespaces - List namespaces`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNamespaces(listNamespacesOptions *ListNamespacesOptions)`, func() {
			listNamespacesOptions := &containerregistryv1.ListNamespacesOptions{}

			result, response, err := containerRegistryService.ListNamespaces(listNamespacesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).To(ContainElement(baseNamespace))

		})
	})

	Describe(`ListNamespaceDetails - Detailed namespace list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNamespaceDetails(listNamespaceDetailsOptions *ListNamespaceDetailsOptions)`, func() {
			listNamespaceDetailsOptions := &containerregistryv1.ListNamespaceDetailsOptions{}

			namespaceDetails, response, err := containerRegistryService.ListNamespaceDetails(listNamespaceDetailsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespaceDetails).ToNot(BeEmpty())

		})
	})

	Describe(`AssignNamespace - Assign namespace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AssignNamespace(assignNamespaceOptions *AssignNamespaceOptions)`, func() {
			assignNamespaceOptions := &containerregistryv1.AssignNamespaceOptions{
				Name:               core.StringPtr(namespaceLink),
				XAuthResourceGroup: core.StringPtr(resouceGroupID),
			}

			namespace, response, err := containerRegistryService.AssignNamespace(assignNamespaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespace).ToNot(BeNil())
		})
	})

	Describe(`GetPlans - Get plans`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPlans(getPlansOptions *GetPlansOptions)`, func() {
			getPlansOptions := &containerregistryv1.GetPlansOptions{}

			plan, response, err := containerRegistryService.GetPlans(getPlansOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(plan).ToNot(BeNil())
		})
	})

	Describe(`UpdatePlans - Update plans`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePlans(updatePlansOptions *UpdatePlansOptions)`, func() {

			Skip("Upgrading plan affects the whole account. Not safe to attempt in this context, skipping tests...")

			updatePlansOptions := &containerregistryv1.UpdatePlansOptions{
				Plan: core.StringPtr("Standard"),
			}

			response, err := containerRegistryService.UpdatePlans(updatePlansOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`UpdateQuota - Update quotas`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateQuota(updateQuotaOptions *UpdateQuotaOptions)`, func() {
			updateQuotaOptions := &containerregistryv1.UpdateQuotaOptions{
				StorageMegabytes: core.Int64Ptr(int64(500)),
				TrafficMegabytes: core.Int64Ptr(int64(4900)),
			}

			response, err := containerRegistryService.UpdateQuota(updateQuotaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`GetQuota - Get quotas`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetQuota(getQuotaOptions *GetQuotaOptions)`, func() {
			getQuotaOptions := &containerregistryv1.GetQuotaOptions{}

			quota, response, err := containerRegistryService.GetQuota(getQuotaOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(*quota.Limit.StorageBytes).To(Equal(int64(524288000)))

		})
	})
	Describe(`ListRetentionPolicies - List retention policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRetentionPolicies(listRetentionPoliciesOptions *ListRetentionPoliciesOptions)`, func() {
			listRetentionPoliciesOptions := &containerregistryv1.ListRetentionPoliciesOptions{}

			mapStringRetentionPolicy, response, err := containerRegistryService.ListRetentionPolicies(listRetentionPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringRetentionPolicy).ToNot(BeNil())
		})
	})

	Describe(`SetRetentionPolicy - Set retention policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetRetentionPolicy(setRetentionPolicyOptions *SetRetentionPolicyOptions)`, func() {
			setRetentionPolicyOptions := &containerregistryv1.SetRetentionPolicyOptions{
				ImagesPerRepo:  core.Int64Ptr(int64(10)),
				Namespace:      core.StringPtr(namespaceLink),
				RetainUntagged: core.BoolPtr(false),
			}

			response, err := containerRegistryService.SetRetentionPolicy(setRetentionPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
	})

	Describe(`AnalyzeRetentionPolicy - Analyze retention policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AnalyzeRetentionPolicy(analyzeRetentionPolicyOptions *AnalyzeRetentionPolicyOptions)`, func() {
			analyzeRetentionPolicyOptions := &containerregistryv1.AnalyzeRetentionPolicyOptions{
				ImagesPerRepo:  core.Int64Ptr(int64(10)),
				Namespace:      core.StringPtr(namespaceLink),
				RetainUntagged: core.BoolPtr(false),
			}

			mapStringstring, response, err := containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringstring).ToNot(BeNil())
		})
	})

	Describe(`GetRetentionPolicy - Get retention policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRetentionPolicy(getRetentionPolicyOptions *GetRetentionPolicyOptions)`, func() {
			getRetentionPolicyOptions := &containerregistryv1.GetRetentionPolicyOptions{
				Namespace: core.StringPtr(namespaceLink),
			}

			retentionPolicy, response, err := containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(*retentionPolicy.ImagesPerRepo).To(Equal(int64(10)))

		})
	})

	Describe(`DeleteImageTag - Delete tag`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteImageTag(deleteImageTagOptions *DeleteImageTagOptions)`, func() {
			deleteImageTagOptions := &containerregistryv1.DeleteImageTagOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)),
			}

			imageDeleteResult, response, err := containerRegistryService.DeleteImageTag(deleteImageTagOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDeleteResult).ToNot(BeNil())
		})
	})
	Describe(`DeleteImage - Delete image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteImage(deleteImageOptions *DeleteImageOptions)`, func() {
			deleteImageOptions := &containerregistryv1.DeleteImageOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest@%s", registryDNSName, namespaceLink, seedDigest)),
			}

			imageDeleteResult, response, err := containerRegistryService.DeleteImage(deleteImageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDeleteResult).ToNot(BeNil())
		})
	})

	Describe(`ListDeletedImages - List deleted images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDeletedImages(listDeletedImagesOptions *ListDeletedImagesOptions)`, func() {
			listDeletedImagesOptions := &containerregistryv1.ListDeletedImagesOptions{
				Namespace: core.StringPtr(namespaceLink),
			}

			mapStringTrash, response, err := containerRegistryService.ListDeletedImages(listDeletedImagesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringTrash[fmt.Sprintf("%s/%s/sdktest@%s", registryDNSName, namespaceLink, seedDigest)]).ToNot(BeNil())

		})
	})

	Describe(`RestoreTags - Restore a digest and all associated tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestoreTags(restoreTagsOptions *RestoreTagsOptions)`, func() {
			restoreTagsOptions := &containerregistryv1.RestoreTagsOptions{
				Digest: core.StringPtr(fmt.Sprintf("%s/%s/sdktest@%s", registryDNSName, namespaceLink, seedDigest)),
			}

			restoreResult, response, err := containerRegistryService.RestoreTags(restoreTagsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(restoreResult).ToNot(BeNil())
		})
	})

	Describe(`RestoreImage - Restore deleted image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestoreImage(restoreImageOptions *RestoreImageOptions)`, func() {
			restoreImageOptions := &containerregistryv1.RestoreImageOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest:nope", registryDNSName, namespaceLink)),
			}

			response, err := containerRegistryService.RestoreImage(restoreImageOptions)
			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`DeleteNamespace - Delete namespace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNamespace(deleteNamespaceOptions *DeleteNamespaceOptions)`, func() {
			deleteNamespaceOptions := &containerregistryv1.DeleteNamespaceOptions{
				Name: core.StringPtr(namespaceLink),
			}

			response, err := containerRegistryService.DeleteNamespace(deleteNamespaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
