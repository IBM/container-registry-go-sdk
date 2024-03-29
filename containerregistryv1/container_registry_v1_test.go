/**
 * (C) Copyright IBM Corp. 2023.
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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ContainerRegistryV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		account := "testString"
		It(`Instantiate service client`, func() {
			containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Account: core.StringPtr(account),
			})
			Expect(containerRegistryService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
				URL: "{BAD_URL_STRING",
				Account: core.StringPtr(account),
			})
			Expect(containerRegistryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
				URL: "https://containerregistryv1/api",
				Account: core.StringPtr(account),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(containerRegistryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{})
			Expect(containerRegistryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		account := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONTAINER_REGISTRY_URL": "https://containerregistryv1/api",
				"CONTAINER_REGISTRY_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1UsingExternalConfig(&containerregistryv1.ContainerRegistryV1Options{
					Account: core.StringPtr(account),
				})
				Expect(containerRegistryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := containerRegistryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != containerRegistryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(containerRegistryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(containerRegistryService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1UsingExternalConfig(&containerregistryv1.ContainerRegistryV1Options{
					URL: "https://testService/api",
					Account: core.StringPtr(account),
				})
				Expect(containerRegistryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := containerRegistryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != containerRegistryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(containerRegistryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(containerRegistryService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1UsingExternalConfig(&containerregistryv1.ContainerRegistryV1Options{
					Account: core.StringPtr(account),
				})
				err := containerRegistryService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := containerRegistryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != containerRegistryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(containerRegistryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(containerRegistryService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONTAINER_REGISTRY_URL": "https://containerregistryv1/api",
				"CONTAINER_REGISTRY_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1UsingExternalConfig(&containerregistryv1.ContainerRegistryV1Options{
				Account: core.StringPtr(account),
			})

			It(`Instantiate service client with error`, func() {
				Expect(containerRegistryService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONTAINER_REGISTRY_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1UsingExternalConfig(&containerregistryv1.ContainerRegistryV1Options{
				URL: "{BAD_URL_STRING",
				Account: core.StringPtr(account),
			})

			It(`Instantiate service client with error`, func() {
				Expect(containerRegistryService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = containerregistryv1.GetServiceURLForRegion("global")
			Expect(url).To(Equal("https://icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://us.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("uk-south")
			Expect(url).To(Equal("https://uk.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://uk.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("eu-central")
			Expect(url).To(Equal("https://de.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://de.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("ap-north")
			Expect(url).To(Equal("https://jp.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("jp-tok")
			Expect(url).To(Equal("https://jp.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("ap-south")
			Expect(url).To(Equal("https://au.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("au-syd")
			Expect(url).To(Equal("https://au.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("jp-osa")
			Expect(url).To(Equal("https://jp2.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("ca-tor")
			Expect(url).To(Equal("https://ca.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("br-sao")
			Expect(url).To(Equal("https://br.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("eu-fr2")
			Expect(url).To(Equal("https://fr2.icr.io"))
			Expect(err).To(BeNil())

			url, err = containerregistryv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAuth(getAuthOptions *GetAuthOptions) - Operation response error`, func() {
		account := "testString"
		getAuthPath := "/api/v1/auth"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAuthPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAuth with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetAuthOptions model
				getAuthOptionsModel := new(containerregistryv1.GetAuthOptions)
				getAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.GetAuth(getAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.GetAuth(getAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAuth(getAuthOptions *GetAuthOptions)`, func() {
		account := "testString"
		getAuthPath := "/api/v1/auth"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAuthPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"iam_authz": true, "private_only": false}`)
				}))
			})
			It(`Invoke GetAuth successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the GetAuthOptions model
				getAuthOptionsModel := new(containerregistryv1.GetAuthOptions)
				getAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.GetAuthWithContext(ctx, getAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.GetAuth(getAuthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.GetAuthWithContext(ctx, getAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAuthPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"iam_authz": true, "private_only": false}`)
				}))
			})
			It(`Invoke GetAuth successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.GetAuth(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAuthOptions model
				getAuthOptionsModel := new(containerregistryv1.GetAuthOptions)
				getAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.GetAuth(getAuthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAuth with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetAuthOptions model
				getAuthOptionsModel := new(containerregistryv1.GetAuthOptions)
				getAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.GetAuth(getAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAuth successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetAuthOptions model
				getAuthOptionsModel := new(containerregistryv1.GetAuthOptions)
				getAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.GetAuth(getAuthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAuth(updateAuthOptions *UpdateAuthOptions)`, func() {
		account := "testString"
		updateAuthPath := "/api/v1/auth"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAuthPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateAuth successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := containerRegistryService.UpdateAuth(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateAuthOptions model
				updateAuthOptionsModel := new(containerregistryv1.UpdateAuthOptions)
				updateAuthOptionsModel.IamAuthz = core.BoolPtr(true)
				updateAuthOptionsModel.PrivateOnly = core.BoolPtr(true)
				updateAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = containerRegistryService.UpdateAuth(updateAuthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateAuth with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the UpdateAuthOptions model
				updateAuthOptionsModel := new(containerregistryv1.UpdateAuthOptions)
				updateAuthOptionsModel.IamAuthz = core.BoolPtr(true)
				updateAuthOptionsModel.PrivateOnly = core.BoolPtr(true)
				updateAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := containerRegistryService.UpdateAuth(updateAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListImages(listImagesOptions *ListImagesOptions) - Operation response error`, func() {
		account := "testString"
		listImagesPath := "/api/v1/images"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listImagesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))
					// TODO: Add check for includeIBM query parameter
					// TODO: Add check for includePrivate query parameter
					// TODO: Add check for includeManifestLists query parameter
					// TODO: Add check for vulnerabilities query parameter
					Expect(req.URL.Query()["repository"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListImages with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(containerregistryv1.ListImagesOptions)
				listImagesOptionsModel.Namespace = core.StringPtr("testString")
				listImagesOptionsModel.IncludeIBM = core.BoolPtr(true)
				listImagesOptionsModel.IncludePrivate = core.BoolPtr(true)
				listImagesOptionsModel.IncludeManifestLists = core.BoolPtr(true)
				listImagesOptionsModel.Vulnerabilities = core.BoolPtr(true)
				listImagesOptionsModel.Repository = core.StringPtr("testString")
				listImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.ListImages(listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.ListImages(listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListImages(listImagesOptions *ListImagesOptions)`, func() {
		account := "testString"
		listImagesPath := "/api/v1/images"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listImagesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))
					// TODO: Add check for includeIBM query parameter
					// TODO: Add check for includePrivate query parameter
					// TODO: Add check for includeManifestLists query parameter
					// TODO: Add check for vulnerabilities query parameter
					Expect(req.URL.Query()["repository"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"ConfigurationIssueCount": 23, "Created": 7, "DigestTags": {"mapKey": ["Inner"]}, "ExemptIssueCount": 16, "Id": "ID", "IssueCount": 10, "Labels": {"mapKey": "Inner"}, "ManifestType": "ManifestType", "ParentId": "ParentID", "RepoDigests": ["RepoDigests"], "RepoTags": ["RepoTags"], "Size": 4, "VirtualSize": 11, "VulnerabilityCount": 18, "Vulnerable": "Vulnerable"}]`)
				}))
			})
			It(`Invoke ListImages successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(containerregistryv1.ListImagesOptions)
				listImagesOptionsModel.Namespace = core.StringPtr("testString")
				listImagesOptionsModel.IncludeIBM = core.BoolPtr(true)
				listImagesOptionsModel.IncludePrivate = core.BoolPtr(true)
				listImagesOptionsModel.IncludeManifestLists = core.BoolPtr(true)
				listImagesOptionsModel.Vulnerabilities = core.BoolPtr(true)
				listImagesOptionsModel.Repository = core.StringPtr("testString")
				listImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.ListImagesWithContext(ctx, listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.ListImages(listImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.ListImagesWithContext(ctx, listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listImagesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))
					// TODO: Add check for includeIBM query parameter
					// TODO: Add check for includePrivate query parameter
					// TODO: Add check for includeManifestLists query parameter
					// TODO: Add check for vulnerabilities query parameter
					Expect(req.URL.Query()["repository"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"ConfigurationIssueCount": 23, "Created": 7, "DigestTags": {"mapKey": ["Inner"]}, "ExemptIssueCount": 16, "Id": "ID", "IssueCount": 10, "Labels": {"mapKey": "Inner"}, "ManifestType": "ManifestType", "ParentId": "ParentID", "RepoDigests": ["RepoDigests"], "RepoTags": ["RepoTags"], "Size": 4, "VirtualSize": 11, "VulnerabilityCount": 18, "Vulnerable": "Vulnerable"}]`)
				}))
			})
			It(`Invoke ListImages successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.ListImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(containerregistryv1.ListImagesOptions)
				listImagesOptionsModel.Namespace = core.StringPtr("testString")
				listImagesOptionsModel.IncludeIBM = core.BoolPtr(true)
				listImagesOptionsModel.IncludePrivate = core.BoolPtr(true)
				listImagesOptionsModel.IncludeManifestLists = core.BoolPtr(true)
				listImagesOptionsModel.Vulnerabilities = core.BoolPtr(true)
				listImagesOptionsModel.Repository = core.StringPtr("testString")
				listImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.ListImages(listImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListImages with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(containerregistryv1.ListImagesOptions)
				listImagesOptionsModel.Namespace = core.StringPtr("testString")
				listImagesOptionsModel.IncludeIBM = core.BoolPtr(true)
				listImagesOptionsModel.IncludePrivate = core.BoolPtr(true)
				listImagesOptionsModel.IncludeManifestLists = core.BoolPtr(true)
				listImagesOptionsModel.Vulnerabilities = core.BoolPtr(true)
				listImagesOptionsModel.Repository = core.StringPtr("testString")
				listImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.ListImages(listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListImages successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(containerregistryv1.ListImagesOptions)
				listImagesOptionsModel.Namespace = core.StringPtr("testString")
				listImagesOptionsModel.IncludeIBM = core.BoolPtr(true)
				listImagesOptionsModel.IncludePrivate = core.BoolPtr(true)
				listImagesOptionsModel.IncludeManifestLists = core.BoolPtr(true)
				listImagesOptionsModel.Vulnerabilities = core.BoolPtr(true)
				listImagesOptionsModel.Repository = core.StringPtr("testString")
				listImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.ListImages(listImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`BulkDeleteImages(bulkDeleteImagesOptions *BulkDeleteImagesOptions) - Operation response error`, func() {
		account := "testString"
		bulkDeleteImagesPath := "/api/v1/images/bulkdelete"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(bulkDeleteImagesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke BulkDeleteImages with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the BulkDeleteImagesOptions model
				bulkDeleteImagesOptionsModel := new(containerregistryv1.BulkDeleteImagesOptions)
				bulkDeleteImagesOptionsModel.BulkDelete = []string{"us.icr.io/birds/woodpecker@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4bbbb", "us.icr.io/birds/bird@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4dddd"}
				bulkDeleteImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`BulkDeleteImages(bulkDeleteImagesOptions *BulkDeleteImagesOptions)`, func() {
		account := "testString"
		bulkDeleteImagesPath := "/api/v1/images/bulkdelete"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(bulkDeleteImagesPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"error": {"mapKey": {"code": "Code", "message": "Message"}}, "success": ["Success"]}`)
				}))
			})
			It(`Invoke BulkDeleteImages successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the BulkDeleteImagesOptions model
				bulkDeleteImagesOptionsModel := new(containerregistryv1.BulkDeleteImagesOptions)
				bulkDeleteImagesOptionsModel.BulkDelete = []string{"us.icr.io/birds/woodpecker@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4bbbb", "us.icr.io/birds/bird@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4dddd"}
				bulkDeleteImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.BulkDeleteImagesWithContext(ctx, bulkDeleteImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.BulkDeleteImagesWithContext(ctx, bulkDeleteImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(bulkDeleteImagesPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"error": {"mapKey": {"code": "Code", "message": "Message"}}, "success": ["Success"]}`)
				}))
			})
			It(`Invoke BulkDeleteImages successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.BulkDeleteImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BulkDeleteImagesOptions model
				bulkDeleteImagesOptionsModel := new(containerregistryv1.BulkDeleteImagesOptions)
				bulkDeleteImagesOptionsModel.BulkDelete = []string{"us.icr.io/birds/woodpecker@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4bbbb", "us.icr.io/birds/bird@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4dddd"}
				bulkDeleteImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke BulkDeleteImages with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the BulkDeleteImagesOptions model
				bulkDeleteImagesOptionsModel := new(containerregistryv1.BulkDeleteImagesOptions)
				bulkDeleteImagesOptionsModel.BulkDelete = []string{"us.icr.io/birds/woodpecker@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4bbbb", "us.icr.io/birds/bird@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4dddd"}
				bulkDeleteImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the BulkDeleteImagesOptions model with no property values
				bulkDeleteImagesOptionsModelNew := new(containerregistryv1.BulkDeleteImagesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke BulkDeleteImages successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the BulkDeleteImagesOptions model
				bulkDeleteImagesOptionsModel := new(containerregistryv1.BulkDeleteImagesOptions)
				bulkDeleteImagesOptionsModel.BulkDelete = []string{"us.icr.io/birds/woodpecker@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4bbbb", "us.icr.io/birds/bird@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4dddd"}
				bulkDeleteImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.BulkDeleteImages(bulkDeleteImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListImageDigests(listImageDigestsOptions *ListImageDigestsOptions) - Operation response error`, func() {
		account := "testString"
		listImageDigestsPath := "/api/v1/images/digests"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listImageDigestsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListImageDigests with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListImageDigestsOptions model
				listImageDigestsOptionsModel := new(containerregistryv1.ListImageDigestsOptions)
				listImageDigestsOptionsModel.ExcludeTagged = core.BoolPtr(false)
				listImageDigestsOptionsModel.ExcludeVa = core.BoolPtr(false)
				listImageDigestsOptionsModel.IncludeIBM = core.BoolPtr(false)
				listImageDigestsOptionsModel.Repositories = []string{"testString"}
				listImageDigestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.ListImageDigests(listImageDigestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.ListImageDigests(listImageDigestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListImageDigests(listImageDigestsOptions *ListImageDigestsOptions)`, func() {
		account := "testString"
		listImageDigestsPath := "/api/v1/images/digests"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listImageDigestsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"created": 7, "id": "ID", "manifestType": "ManifestType", "repoTags": {"anyKey": "anyValue"}, "size": 4}]`)
				}))
			})
			It(`Invoke ListImageDigests successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the ListImageDigestsOptions model
				listImageDigestsOptionsModel := new(containerregistryv1.ListImageDigestsOptions)
				listImageDigestsOptionsModel.ExcludeTagged = core.BoolPtr(false)
				listImageDigestsOptionsModel.ExcludeVa = core.BoolPtr(false)
				listImageDigestsOptionsModel.IncludeIBM = core.BoolPtr(false)
				listImageDigestsOptionsModel.Repositories = []string{"testString"}
				listImageDigestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.ListImageDigestsWithContext(ctx, listImageDigestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.ListImageDigests(listImageDigestsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.ListImageDigestsWithContext(ctx, listImageDigestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listImageDigestsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"created": 7, "id": "ID", "manifestType": "ManifestType", "repoTags": {"anyKey": "anyValue"}, "size": 4}]`)
				}))
			})
			It(`Invoke ListImageDigests successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.ListImageDigests(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListImageDigestsOptions model
				listImageDigestsOptionsModel := new(containerregistryv1.ListImageDigestsOptions)
				listImageDigestsOptionsModel.ExcludeTagged = core.BoolPtr(false)
				listImageDigestsOptionsModel.ExcludeVa = core.BoolPtr(false)
				listImageDigestsOptionsModel.IncludeIBM = core.BoolPtr(false)
				listImageDigestsOptionsModel.Repositories = []string{"testString"}
				listImageDigestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.ListImageDigests(listImageDigestsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListImageDigests with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListImageDigestsOptions model
				listImageDigestsOptionsModel := new(containerregistryv1.ListImageDigestsOptions)
				listImageDigestsOptionsModel.ExcludeTagged = core.BoolPtr(false)
				listImageDigestsOptionsModel.ExcludeVa = core.BoolPtr(false)
				listImageDigestsOptionsModel.IncludeIBM = core.BoolPtr(false)
				listImageDigestsOptionsModel.Repositories = []string{"testString"}
				listImageDigestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.ListImageDigests(listImageDigestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListImageDigests successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListImageDigestsOptions model
				listImageDigestsOptionsModel := new(containerregistryv1.ListImageDigestsOptions)
				listImageDigestsOptionsModel.ExcludeTagged = core.BoolPtr(false)
				listImageDigestsOptionsModel.ExcludeVa = core.BoolPtr(false)
				listImageDigestsOptionsModel.IncludeIBM = core.BoolPtr(false)
				listImageDigestsOptionsModel.Repositories = []string{"testString"}
				listImageDigestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.ListImageDigests(listImageDigestsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TagImage(tagImageOptions *TagImageOptions)`, func() {
		account := "testString"
		tagImagePath := "/api/v1/images/tags"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(tagImagePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["fromimage"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["toimage"]).To(Equal([]string{"testString"}))
					res.WriteHeader(201)
				}))
			})
			It(`Invoke TagImage successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := containerRegistryService.TagImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the TagImageOptions model
				tagImageOptionsModel := new(containerregistryv1.TagImageOptions)
				tagImageOptionsModel.Fromimage = core.StringPtr("testString")
				tagImageOptionsModel.Toimage = core.StringPtr("testString")
				tagImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = containerRegistryService.TagImage(tagImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke TagImage with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the TagImageOptions model
				tagImageOptionsModel := new(containerregistryv1.TagImageOptions)
				tagImageOptionsModel.Fromimage = core.StringPtr("testString")
				tagImageOptionsModel.Toimage = core.StringPtr("testString")
				tagImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := containerRegistryService.TagImage(tagImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the TagImageOptions model with no property values
				tagImageOptionsModelNew := new(containerregistryv1.TagImageOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = containerRegistryService.TagImage(tagImageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteImage(deleteImageOptions *DeleteImageOptions) - Operation response error`, func() {
		account := "testString"
		deleteImagePath := "/api/v1/images/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteImagePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteImage with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the DeleteImageOptions model
				deleteImageOptionsModel := new(containerregistryv1.DeleteImageOptions)
				deleteImageOptionsModel.Image = core.StringPtr("testString")
				deleteImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteImage(deleteImageOptions *DeleteImageOptions)`, func() {
		account := "testString"
		deleteImagePath := "/api/v1/images/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteImagePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Untagged": "Untagged"}`)
				}))
			})
			It(`Invoke DeleteImage successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteImageOptions model
				deleteImageOptionsModel := new(containerregistryv1.DeleteImageOptions)
				deleteImageOptionsModel.Image = core.StringPtr("testString")
				deleteImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.DeleteImageWithContext(ctx, deleteImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.DeleteImageWithContext(ctx, deleteImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteImagePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Untagged": "Untagged"}`)
				}))
			})
			It(`Invoke DeleteImage successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.DeleteImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteImageOptions model
				deleteImageOptionsModel := new(containerregistryv1.DeleteImageOptions)
				deleteImageOptionsModel.Image = core.StringPtr("testString")
				deleteImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteImage with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the DeleteImageOptions model
				deleteImageOptionsModel := new(containerregistryv1.DeleteImageOptions)
				deleteImageOptionsModel.Image = core.StringPtr("testString")
				deleteImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteImageOptions model with no property values
				deleteImageOptionsModelNew := new(containerregistryv1.DeleteImageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.DeleteImage(deleteImageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteImage successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the DeleteImageOptions model
				deleteImageOptionsModel := new(containerregistryv1.DeleteImageOptions)
				deleteImageOptionsModel.Image = core.StringPtr("testString")
				deleteImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`InspectImage(inspectImageOptions *InspectImageOptions) - Operation response error`, func() {
		account := "testString"
		inspectImagePath := "/api/v1/images/testString/json"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(inspectImagePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke InspectImage with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the InspectImageOptions model
				inspectImageOptionsModel := new(containerregistryv1.InspectImageOptions)
				inspectImageOptionsModel.Image = core.StringPtr("testString")
				inspectImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.InspectImage(inspectImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.InspectImage(inspectImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`InspectImage(inspectImageOptions *InspectImageOptions)`, func() {
		account := "testString"
		inspectImagePath := "/api/v1/images/testString/json"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(inspectImagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Architecture": "Architecture", "Author": "Author", "Comment": "Comment", "Config": {"ArgsEscaped": false, "AttachStderr": true, "AttachStdin": false, "AttachStdout": true, "Cmd": ["Cmd"], "Domainname": "Domainname", "Entrypoint": ["Entrypoint"], "Env": ["Env"], "ExposedPorts": {"anyKey": "anyValue"}, "Healthcheck": {"Interval": 8, "Retries": 7, "Test": ["Test"], "Timeout": 7}, "Hostname": "Hostname", "Image": "Image", "Labels": {"mapKey": "Inner"}, "MacAddress": "MacAddress", "NetworkDisabled": false, "OnBuild": ["OnBuild"], "OpenStdin": false, "Shell": ["Shell"], "StdinOnce": false, "StopSignal": "StopSignal", "StopTimeout": 11, "Tty": false, "User": "User", "Volumes": {"anyKey": "anyValue"}, "WorkingDir": "WorkingDir"}, "Container": "Container", "ContainerConfig": {"ArgsEscaped": false, "AttachStderr": true, "AttachStdin": false, "AttachStdout": true, "Cmd": ["Cmd"], "Domainname": "Domainname", "Entrypoint": ["Entrypoint"], "Env": ["Env"], "ExposedPorts": {"anyKey": "anyValue"}, "Healthcheck": {"Interval": 8, "Retries": 7, "Test": ["Test"], "Timeout": 7}, "Hostname": "Hostname", "Image": "Image", "Labels": {"mapKey": "Inner"}, "MacAddress": "MacAddress", "NetworkDisabled": false, "OnBuild": ["OnBuild"], "OpenStdin": false, "Shell": ["Shell"], "StdinOnce": false, "StopSignal": "StopSignal", "StopTimeout": 11, "Tty": false, "User": "User", "Volumes": {"anyKey": "anyValue"}, "WorkingDir": "WorkingDir"}, "Created": "Created", "DockerVersion": "DockerVersion", "Id": "ID", "ManifestType": "ManifestType", "Os": "Os", "OsVersion": "OsVersion", "Parent": "Parent", "RootFS": {"BaseLayer": "BaseLayer", "Layers": ["Layers"], "Type": "Type"}, "Size": 4, "VirtualSize": 11}`)
				}))
			})
			It(`Invoke InspectImage successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the InspectImageOptions model
				inspectImageOptionsModel := new(containerregistryv1.InspectImageOptions)
				inspectImageOptionsModel.Image = core.StringPtr("testString")
				inspectImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.InspectImageWithContext(ctx, inspectImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.InspectImage(inspectImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.InspectImageWithContext(ctx, inspectImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(inspectImagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Architecture": "Architecture", "Author": "Author", "Comment": "Comment", "Config": {"ArgsEscaped": false, "AttachStderr": true, "AttachStdin": false, "AttachStdout": true, "Cmd": ["Cmd"], "Domainname": "Domainname", "Entrypoint": ["Entrypoint"], "Env": ["Env"], "ExposedPorts": {"anyKey": "anyValue"}, "Healthcheck": {"Interval": 8, "Retries": 7, "Test": ["Test"], "Timeout": 7}, "Hostname": "Hostname", "Image": "Image", "Labels": {"mapKey": "Inner"}, "MacAddress": "MacAddress", "NetworkDisabled": false, "OnBuild": ["OnBuild"], "OpenStdin": false, "Shell": ["Shell"], "StdinOnce": false, "StopSignal": "StopSignal", "StopTimeout": 11, "Tty": false, "User": "User", "Volumes": {"anyKey": "anyValue"}, "WorkingDir": "WorkingDir"}, "Container": "Container", "ContainerConfig": {"ArgsEscaped": false, "AttachStderr": true, "AttachStdin": false, "AttachStdout": true, "Cmd": ["Cmd"], "Domainname": "Domainname", "Entrypoint": ["Entrypoint"], "Env": ["Env"], "ExposedPorts": {"anyKey": "anyValue"}, "Healthcheck": {"Interval": 8, "Retries": 7, "Test": ["Test"], "Timeout": 7}, "Hostname": "Hostname", "Image": "Image", "Labels": {"mapKey": "Inner"}, "MacAddress": "MacAddress", "NetworkDisabled": false, "OnBuild": ["OnBuild"], "OpenStdin": false, "Shell": ["Shell"], "StdinOnce": false, "StopSignal": "StopSignal", "StopTimeout": 11, "Tty": false, "User": "User", "Volumes": {"anyKey": "anyValue"}, "WorkingDir": "WorkingDir"}, "Created": "Created", "DockerVersion": "DockerVersion", "Id": "ID", "ManifestType": "ManifestType", "Os": "Os", "OsVersion": "OsVersion", "Parent": "Parent", "RootFS": {"BaseLayer": "BaseLayer", "Layers": ["Layers"], "Type": "Type"}, "Size": 4, "VirtualSize": 11}`)
				}))
			})
			It(`Invoke InspectImage successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.InspectImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InspectImageOptions model
				inspectImageOptionsModel := new(containerregistryv1.InspectImageOptions)
				inspectImageOptionsModel.Image = core.StringPtr("testString")
				inspectImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.InspectImage(inspectImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke InspectImage with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the InspectImageOptions model
				inspectImageOptionsModel := new(containerregistryv1.InspectImageOptions)
				inspectImageOptionsModel.Image = core.StringPtr("testString")
				inspectImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.InspectImage(inspectImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the InspectImageOptions model with no property values
				inspectImageOptionsModelNew := new(containerregistryv1.InspectImageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.InspectImage(inspectImageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke InspectImage successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the InspectImageOptions model
				inspectImageOptionsModel := new(containerregistryv1.InspectImageOptions)
				inspectImageOptionsModel.Image = core.StringPtr("testString")
				inspectImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.InspectImage(inspectImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetImageManifest(getImageManifestOptions *GetImageManifestOptions)`, func() {
		account := "testString"
		getImageManifestPath := "/api/v1/images/testString/manifest"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getImageManifestPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"anyKey": "anyValue"}`)
				}))
			})
			It(`Invoke GetImageManifest successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the GetImageManifestOptions model
				getImageManifestOptionsModel := new(containerregistryv1.GetImageManifestOptions)
				getImageManifestOptionsModel.Image = core.StringPtr("testString")
				getImageManifestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.GetImageManifestWithContext(ctx, getImageManifestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.GetImageManifest(getImageManifestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.GetImageManifestWithContext(ctx, getImageManifestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getImageManifestPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"anyKey": "anyValue"}`)
				}))
			})
			It(`Invoke GetImageManifest successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.GetImageManifest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetImageManifestOptions model
				getImageManifestOptionsModel := new(containerregistryv1.GetImageManifestOptions)
				getImageManifestOptionsModel.Image = core.StringPtr("testString")
				getImageManifestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.GetImageManifest(getImageManifestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetImageManifest with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetImageManifestOptions model
				getImageManifestOptionsModel := new(containerregistryv1.GetImageManifestOptions)
				getImageManifestOptionsModel.Image = core.StringPtr("testString")
				getImageManifestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.GetImageManifest(getImageManifestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetImageManifestOptions model with no property values
				getImageManifestOptionsModelNew := new(containerregistryv1.GetImageManifestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.GetImageManifest(getImageManifestOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetImageManifest successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetImageManifestOptions model
				getImageManifestOptionsModel := new(containerregistryv1.GetImageManifestOptions)
				getImageManifestOptionsModel.Image = core.StringPtr("testString")
				getImageManifestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.GetImageManifest(getImageManifestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMessages(getMessagesOptions *GetMessagesOptions)`, func() {
		account := "testString"
		getMessagesPath := "/api/v1/messages"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMessagesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"Hello, world!"`)
				}))
			})
			It(`Invoke GetMessages successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the GetMessagesOptions model
				getMessagesOptionsModel := new(containerregistryv1.GetMessagesOptions)
				getMessagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.GetMessagesWithContext(ctx, getMessagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.GetMessages(getMessagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.GetMessagesWithContext(ctx, getMessagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMessagesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"Hello, world!"`)
				}))
			})
			It(`Invoke GetMessages successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.GetMessages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMessagesOptions model
				getMessagesOptionsModel := new(containerregistryv1.GetMessagesOptions)
				getMessagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.GetMessages(getMessagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMessages with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetMessagesOptions model
				getMessagesOptionsModel := new(containerregistryv1.GetMessagesOptions)
				getMessagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.GetMessages(getMessagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetMessages successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetMessagesOptions model
				getMessagesOptionsModel := new(containerregistryv1.GetMessagesOptions)
				getMessagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.GetMessages(getMessagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNamespaces(listNamespacesOptions *ListNamespacesOptions)`, func() {
		account := "testString"
		listNamespacesPath := "/api/v1/namespaces"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNamespacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `["OperationResponse"]`)
				}))
			})
			It(`Invoke ListNamespaces successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the ListNamespacesOptions model
				listNamespacesOptionsModel := new(containerregistryv1.ListNamespacesOptions)
				listNamespacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.ListNamespacesWithContext(ctx, listNamespacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.ListNamespaces(listNamespacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.ListNamespacesWithContext(ctx, listNamespacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNamespacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `["OperationResponse"]`)
				}))
			})
			It(`Invoke ListNamespaces successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.ListNamespaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNamespacesOptions model
				listNamespacesOptionsModel := new(containerregistryv1.ListNamespacesOptions)
				listNamespacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.ListNamespaces(listNamespacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListNamespaces with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListNamespacesOptions model
				listNamespacesOptionsModel := new(containerregistryv1.ListNamespacesOptions)
				listNamespacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.ListNamespaces(listNamespacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListNamespaces successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListNamespacesOptions model
				listNamespacesOptionsModel := new(containerregistryv1.ListNamespacesOptions)
				listNamespacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.ListNamespaces(listNamespacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNamespaceDetails(listNamespaceDetailsOptions *ListNamespaceDetailsOptions) - Operation response error`, func() {
		account := "testString"
		listNamespaceDetailsPath := "/api/v1/namespaces/details"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNamespaceDetailsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNamespaceDetails with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListNamespaceDetailsOptions model
				listNamespaceDetailsOptionsModel := new(containerregistryv1.ListNamespaceDetailsOptions)
				listNamespaceDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.ListNamespaceDetails(listNamespaceDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.ListNamespaceDetails(listNamespaceDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNamespaceDetails(listNamespaceDetailsOptions *ListNamespaceDetailsOptions)`, func() {
		account := "testString"
		listNamespaceDetailsPath := "/api/v1/namespaces/details"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNamespaceDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"account": "Account", "created_date": "CreatedDate", "crn": "CRN", "name": "Name", "resource_created_date": "ResourceCreatedDate", "resource_group": "ResourceGroup", "updated_date": "UpdatedDate"}]`)
				}))
			})
			It(`Invoke ListNamespaceDetails successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the ListNamespaceDetailsOptions model
				listNamespaceDetailsOptionsModel := new(containerregistryv1.ListNamespaceDetailsOptions)
				listNamespaceDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.ListNamespaceDetailsWithContext(ctx, listNamespaceDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.ListNamespaceDetails(listNamespaceDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.ListNamespaceDetailsWithContext(ctx, listNamespaceDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNamespaceDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"account": "Account", "created_date": "CreatedDate", "crn": "CRN", "name": "Name", "resource_created_date": "ResourceCreatedDate", "resource_group": "ResourceGroup", "updated_date": "UpdatedDate"}]`)
				}))
			})
			It(`Invoke ListNamespaceDetails successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.ListNamespaceDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNamespaceDetailsOptions model
				listNamespaceDetailsOptionsModel := new(containerregistryv1.ListNamespaceDetailsOptions)
				listNamespaceDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.ListNamespaceDetails(listNamespaceDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListNamespaceDetails with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListNamespaceDetailsOptions model
				listNamespaceDetailsOptionsModel := new(containerregistryv1.ListNamespaceDetailsOptions)
				listNamespaceDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.ListNamespaceDetails(listNamespaceDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListNamespaceDetails successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListNamespaceDetailsOptions model
				listNamespaceDetailsOptionsModel := new(containerregistryv1.ListNamespaceDetailsOptions)
				listNamespaceDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.ListNamespaceDetails(listNamespaceDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateNamespace(createNamespaceOptions *CreateNamespaceOptions) - Operation response error`, func() {
		account := "testString"
		createNamespacePath := "/api/v1/namespaces/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createNamespacePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Auth-Resource-Group"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Resource-Group"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateNamespace with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the CreateNamespaceOptions model
				createNamespaceOptionsModel := new(containerregistryv1.CreateNamespaceOptions)
				createNamespaceOptionsModel.Name = core.StringPtr("testString")
				createNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				createNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.CreateNamespace(createNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.CreateNamespace(createNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateNamespace(createNamespaceOptions *CreateNamespaceOptions)`, func() {
		account := "testString"
		createNamespacePath := "/api/v1/namespaces/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createNamespacePath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Auth-Resource-Group"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Resource-Group"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"namespace": "Namespace"}`)
				}))
			})
			It(`Invoke CreateNamespace successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the CreateNamespaceOptions model
				createNamespaceOptionsModel := new(containerregistryv1.CreateNamespaceOptions)
				createNamespaceOptionsModel.Name = core.StringPtr("testString")
				createNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				createNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.CreateNamespaceWithContext(ctx, createNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.CreateNamespace(createNamespaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.CreateNamespaceWithContext(ctx, createNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createNamespacePath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Auth-Resource-Group"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Resource-Group"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"namespace": "Namespace"}`)
				}))
			})
			It(`Invoke CreateNamespace successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.CreateNamespace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateNamespaceOptions model
				createNamespaceOptionsModel := new(containerregistryv1.CreateNamespaceOptions)
				createNamespaceOptionsModel.Name = core.StringPtr("testString")
				createNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				createNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.CreateNamespace(createNamespaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateNamespace with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the CreateNamespaceOptions model
				createNamespaceOptionsModel := new(containerregistryv1.CreateNamespaceOptions)
				createNamespaceOptionsModel.Name = core.StringPtr("testString")
				createNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				createNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.CreateNamespace(createNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateNamespaceOptions model with no property values
				createNamespaceOptionsModelNew := new(containerregistryv1.CreateNamespaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.CreateNamespace(createNamespaceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateNamespace successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the CreateNamespaceOptions model
				createNamespaceOptionsModel := new(containerregistryv1.CreateNamespaceOptions)
				createNamespaceOptionsModel.Name = core.StringPtr("testString")
				createNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				createNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.CreateNamespace(createNamespaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AssignNamespace(assignNamespaceOptions *AssignNamespaceOptions) - Operation response error`, func() {
		account := "testString"
		assignNamespacePath := "/api/v1/namespaces/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(assignNamespacePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Auth-Resource-Group"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Resource-Group"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AssignNamespace with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the AssignNamespaceOptions model
				assignNamespaceOptionsModel := new(containerregistryv1.AssignNamespaceOptions)
				assignNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				assignNamespaceOptionsModel.Name = core.StringPtr("testString")
				assignNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.AssignNamespace(assignNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.AssignNamespace(assignNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AssignNamespace(assignNamespaceOptions *AssignNamespaceOptions)`, func() {
		account := "testString"
		assignNamespacePath := "/api/v1/namespaces/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(assignNamespacePath))
					Expect(req.Method).To(Equal("PATCH"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Auth-Resource-Group"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Resource-Group"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"namespace": "Namespace"}`)
				}))
			})
			It(`Invoke AssignNamespace successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the AssignNamespaceOptions model
				assignNamespaceOptionsModel := new(containerregistryv1.AssignNamespaceOptions)
				assignNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				assignNamespaceOptionsModel.Name = core.StringPtr("testString")
				assignNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.AssignNamespaceWithContext(ctx, assignNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.AssignNamespace(assignNamespaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.AssignNamespaceWithContext(ctx, assignNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(assignNamespacePath))
					Expect(req.Method).To(Equal("PATCH"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Auth-Resource-Group"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Resource-Group"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"namespace": "Namespace"}`)
				}))
			})
			It(`Invoke AssignNamespace successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.AssignNamespace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AssignNamespaceOptions model
				assignNamespaceOptionsModel := new(containerregistryv1.AssignNamespaceOptions)
				assignNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				assignNamespaceOptionsModel.Name = core.StringPtr("testString")
				assignNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.AssignNamespace(assignNamespaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AssignNamespace with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the AssignNamespaceOptions model
				assignNamespaceOptionsModel := new(containerregistryv1.AssignNamespaceOptions)
				assignNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				assignNamespaceOptionsModel.Name = core.StringPtr("testString")
				assignNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.AssignNamespace(assignNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AssignNamespaceOptions model with no property values
				assignNamespaceOptionsModelNew := new(containerregistryv1.AssignNamespaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.AssignNamespace(assignNamespaceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke AssignNamespace successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the AssignNamespaceOptions model
				assignNamespaceOptionsModel := new(containerregistryv1.AssignNamespaceOptions)
				assignNamespaceOptionsModel.XAuthResourceGroup = core.StringPtr("testString")
				assignNamespaceOptionsModel.Name = core.StringPtr("testString")
				assignNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.AssignNamespace(assignNamespaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteNamespace(deleteNamespaceOptions *DeleteNamespaceOptions)`, func() {
		account := "testString"
		deleteNamespacePath := "/api/v1/namespaces/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteNamespacePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteNamespace successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := containerRegistryService.DeleteNamespace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNamespaceOptions model
				deleteNamespaceOptionsModel := new(containerregistryv1.DeleteNamespaceOptions)
				deleteNamespaceOptionsModel.Name = core.StringPtr("testString")
				deleteNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = containerRegistryService.DeleteNamespace(deleteNamespaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteNamespace with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the DeleteNamespaceOptions model
				deleteNamespaceOptionsModel := new(containerregistryv1.DeleteNamespaceOptions)
				deleteNamespaceOptionsModel.Name = core.StringPtr("testString")
				deleteNamespaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := containerRegistryService.DeleteNamespace(deleteNamespaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteNamespaceOptions model with no property values
				deleteNamespaceOptionsModelNew := new(containerregistryv1.DeleteNamespaceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = containerRegistryService.DeleteNamespace(deleteNamespaceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPlans(getPlansOptions *GetPlansOptions) - Operation response error`, func() {
		account := "testString"
		getPlansPath := "/api/v1/plans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPlansPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPlans with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetPlansOptions model
				getPlansOptionsModel := new(containerregistryv1.GetPlansOptions)
				getPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.GetPlans(getPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.GetPlans(getPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPlans(getPlansOptions *GetPlansOptions)`, func() {
		account := "testString"
		getPlansPath := "/api/v1/plans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPlansPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plan": "Plan"}`)
				}))
			})
			It(`Invoke GetPlans successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the GetPlansOptions model
				getPlansOptionsModel := new(containerregistryv1.GetPlansOptions)
				getPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.GetPlansWithContext(ctx, getPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.GetPlans(getPlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.GetPlansWithContext(ctx, getPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPlansPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plan": "Plan"}`)
				}))
			})
			It(`Invoke GetPlans successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.GetPlans(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPlansOptions model
				getPlansOptionsModel := new(containerregistryv1.GetPlansOptions)
				getPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.GetPlans(getPlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPlans with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetPlansOptions model
				getPlansOptionsModel := new(containerregistryv1.GetPlansOptions)
				getPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.GetPlans(getPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetPlans successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetPlansOptions model
				getPlansOptionsModel := new(containerregistryv1.GetPlansOptions)
				getPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.GetPlans(getPlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePlans(updatePlansOptions *UpdatePlansOptions)`, func() {
		account := "testString"
		updatePlansPath := "/api/v1/plans"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePlansPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdatePlans successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := containerRegistryService.UpdatePlans(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdatePlansOptions model
				updatePlansOptionsModel := new(containerregistryv1.UpdatePlansOptions)
				updatePlansOptionsModel.Plan = core.StringPtr("Standard")
				updatePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = containerRegistryService.UpdatePlans(updatePlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdatePlans with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the UpdatePlansOptions model
				updatePlansOptionsModel := new(containerregistryv1.UpdatePlansOptions)
				updatePlansOptionsModel.Plan = core.StringPtr("Standard")
				updatePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := containerRegistryService.UpdatePlans(updatePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetQuota(getQuotaOptions *GetQuotaOptions) - Operation response error`, func() {
		account := "testString"
		getQuotaPath := "/api/v1/quotas"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getQuotaPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetQuota with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetQuotaOptions model
				getQuotaOptionsModel := new(containerregistryv1.GetQuotaOptions)
				getQuotaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.GetQuota(getQuotaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.GetQuota(getQuotaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetQuota(getQuotaOptions *GetQuotaOptions)`, func() {
		account := "testString"
		getQuotaPath := "/api/v1/quotas"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getQuotaPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": {"storage_bytes": 12, "traffic_bytes": 12}, "usage": {"storage_bytes": 12, "traffic_bytes": 12}}`)
				}))
			})
			It(`Invoke GetQuota successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the GetQuotaOptions model
				getQuotaOptionsModel := new(containerregistryv1.GetQuotaOptions)
				getQuotaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.GetQuotaWithContext(ctx, getQuotaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.GetQuota(getQuotaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.GetQuotaWithContext(ctx, getQuotaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getQuotaPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": {"storage_bytes": 12, "traffic_bytes": 12}, "usage": {"storage_bytes": 12, "traffic_bytes": 12}}`)
				}))
			})
			It(`Invoke GetQuota successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.GetQuota(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetQuotaOptions model
				getQuotaOptionsModel := new(containerregistryv1.GetQuotaOptions)
				getQuotaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.GetQuota(getQuotaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetQuota with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetQuotaOptions model
				getQuotaOptionsModel := new(containerregistryv1.GetQuotaOptions)
				getQuotaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.GetQuota(getQuotaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetQuota successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetQuotaOptions model
				getQuotaOptionsModel := new(containerregistryv1.GetQuotaOptions)
				getQuotaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.GetQuota(getQuotaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateQuota(updateQuotaOptions *UpdateQuotaOptions)`, func() {
		account := "testString"
		updateQuotaPath := "/api/v1/quotas"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateQuotaPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateQuota successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := containerRegistryService.UpdateQuota(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateQuotaOptions model
				updateQuotaOptionsModel := new(containerregistryv1.UpdateQuotaOptions)
				updateQuotaOptionsModel.StorageMegabytes = core.Int64Ptr(int64(26))
				updateQuotaOptionsModel.TrafficMegabytes = core.Int64Ptr(int64(480))
				updateQuotaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = containerRegistryService.UpdateQuota(updateQuotaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateQuota with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the UpdateQuotaOptions model
				updateQuotaOptionsModel := new(containerregistryv1.UpdateQuotaOptions)
				updateQuotaOptionsModel.StorageMegabytes = core.Int64Ptr(int64(26))
				updateQuotaOptionsModel.TrafficMegabytes = core.Int64Ptr(int64(480))
				updateQuotaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := containerRegistryService.UpdateQuota(updateQuotaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRetentionPolicies(listRetentionPoliciesOptions *ListRetentionPoliciesOptions) - Operation response error`, func() {
		account := "testString"
		listRetentionPoliciesPath := "/api/v1/retentions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRetentionPoliciesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRetentionPolicies with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListRetentionPoliciesOptions model
				listRetentionPoliciesOptionsModel := new(containerregistryv1.ListRetentionPoliciesOptions)
				listRetentionPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.ListRetentionPolicies(listRetentionPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.ListRetentionPolicies(listRetentionPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRetentionPolicies(listRetentionPoliciesOptions *ListRetentionPoliciesOptions)`, func() {
		account := "testString"
		listRetentionPoliciesPath := "/api/v1/retentions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRetentionPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": {"images_per_repo": 13, "namespace": "Namespace", "retain_untagged": true}}`)
				}))
			})
			It(`Invoke ListRetentionPolicies successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the ListRetentionPoliciesOptions model
				listRetentionPoliciesOptionsModel := new(containerregistryv1.ListRetentionPoliciesOptions)
				listRetentionPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.ListRetentionPoliciesWithContext(ctx, listRetentionPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.ListRetentionPolicies(listRetentionPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.ListRetentionPoliciesWithContext(ctx, listRetentionPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRetentionPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": {"images_per_repo": 13, "namespace": "Namespace", "retain_untagged": true}}`)
				}))
			})
			It(`Invoke ListRetentionPolicies successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.ListRetentionPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRetentionPoliciesOptions model
				listRetentionPoliciesOptionsModel := new(containerregistryv1.ListRetentionPoliciesOptions)
				listRetentionPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.ListRetentionPolicies(listRetentionPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRetentionPolicies with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListRetentionPoliciesOptions model
				listRetentionPoliciesOptionsModel := new(containerregistryv1.ListRetentionPoliciesOptions)
				listRetentionPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.ListRetentionPolicies(listRetentionPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListRetentionPolicies successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListRetentionPoliciesOptions model
				listRetentionPoliciesOptionsModel := new(containerregistryv1.ListRetentionPoliciesOptions)
				listRetentionPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.ListRetentionPolicies(listRetentionPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetRetentionPolicy(setRetentionPolicyOptions *SetRetentionPolicyOptions)`, func() {
		account := "testString"
		setRetentionPolicyPath := "/api/v1/retentions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setRetentionPolicyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke SetRetentionPolicy successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := containerRegistryService.SetRetentionPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SetRetentionPolicyOptions model
				setRetentionPolicyOptionsModel := new(containerregistryv1.SetRetentionPolicyOptions)
				setRetentionPolicyOptionsModel.Namespace = core.StringPtr("birds")
				setRetentionPolicyOptionsModel.ImagesPerRepo = core.Int64Ptr(int64(10))
				setRetentionPolicyOptionsModel.RetainUntagged = core.BoolPtr(false)
				setRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = containerRegistryService.SetRetentionPolicy(setRetentionPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SetRetentionPolicy with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the SetRetentionPolicyOptions model
				setRetentionPolicyOptionsModel := new(containerregistryv1.SetRetentionPolicyOptions)
				setRetentionPolicyOptionsModel.Namespace = core.StringPtr("birds")
				setRetentionPolicyOptionsModel.ImagesPerRepo = core.Int64Ptr(int64(10))
				setRetentionPolicyOptionsModel.RetainUntagged = core.BoolPtr(false)
				setRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := containerRegistryService.SetRetentionPolicy(setRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SetRetentionPolicyOptions model with no property values
				setRetentionPolicyOptionsModelNew := new(containerregistryv1.SetRetentionPolicyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = containerRegistryService.SetRetentionPolicy(setRetentionPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AnalyzeRetentionPolicy(analyzeRetentionPolicyOptions *AnalyzeRetentionPolicyOptions) - Operation response error`, func() {
		account := "testString"
		analyzeRetentionPolicyPath := "/api/v1/retentions/analyze"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzeRetentionPolicyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AnalyzeRetentionPolicy with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the AnalyzeRetentionPolicyOptions model
				analyzeRetentionPolicyOptionsModel := new(containerregistryv1.AnalyzeRetentionPolicyOptions)
				analyzeRetentionPolicyOptionsModel.Namespace = core.StringPtr("birds")
				analyzeRetentionPolicyOptionsModel.ImagesPerRepo = core.Int64Ptr(int64(10))
				analyzeRetentionPolicyOptionsModel.RetainUntagged = core.BoolPtr(false)
				analyzeRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AnalyzeRetentionPolicy(analyzeRetentionPolicyOptions *AnalyzeRetentionPolicyOptions)`, func() {
		account := "testString"
		analyzeRetentionPolicyPath := "/api/v1/retentions/analyze"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzeRetentionPolicyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": ["Inner"]}`)
				}))
			})
			It(`Invoke AnalyzeRetentionPolicy successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the AnalyzeRetentionPolicyOptions model
				analyzeRetentionPolicyOptionsModel := new(containerregistryv1.AnalyzeRetentionPolicyOptions)
				analyzeRetentionPolicyOptionsModel.Namespace = core.StringPtr("birds")
				analyzeRetentionPolicyOptionsModel.ImagesPerRepo = core.Int64Ptr(int64(10))
				analyzeRetentionPolicyOptionsModel.RetainUntagged = core.BoolPtr(false)
				analyzeRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.AnalyzeRetentionPolicyWithContext(ctx, analyzeRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.AnalyzeRetentionPolicyWithContext(ctx, analyzeRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzeRetentionPolicyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": ["Inner"]}`)
				}))
			})
			It(`Invoke AnalyzeRetentionPolicy successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.AnalyzeRetentionPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AnalyzeRetentionPolicyOptions model
				analyzeRetentionPolicyOptionsModel := new(containerregistryv1.AnalyzeRetentionPolicyOptions)
				analyzeRetentionPolicyOptionsModel.Namespace = core.StringPtr("birds")
				analyzeRetentionPolicyOptionsModel.ImagesPerRepo = core.Int64Ptr(int64(10))
				analyzeRetentionPolicyOptionsModel.RetainUntagged = core.BoolPtr(false)
				analyzeRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AnalyzeRetentionPolicy with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the AnalyzeRetentionPolicyOptions model
				analyzeRetentionPolicyOptionsModel := new(containerregistryv1.AnalyzeRetentionPolicyOptions)
				analyzeRetentionPolicyOptionsModel.Namespace = core.StringPtr("birds")
				analyzeRetentionPolicyOptionsModel.ImagesPerRepo = core.Int64Ptr(int64(10))
				analyzeRetentionPolicyOptionsModel.RetainUntagged = core.BoolPtr(false)
				analyzeRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AnalyzeRetentionPolicyOptions model with no property values
				analyzeRetentionPolicyOptionsModelNew := new(containerregistryv1.AnalyzeRetentionPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke AnalyzeRetentionPolicy successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the AnalyzeRetentionPolicyOptions model
				analyzeRetentionPolicyOptionsModel := new(containerregistryv1.AnalyzeRetentionPolicyOptions)
				analyzeRetentionPolicyOptionsModel.Namespace = core.StringPtr("birds")
				analyzeRetentionPolicyOptionsModel.ImagesPerRepo = core.Int64Ptr(int64(10))
				analyzeRetentionPolicyOptionsModel.RetainUntagged = core.BoolPtr(false)
				analyzeRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRetentionPolicy(getRetentionPolicyOptions *GetRetentionPolicyOptions) - Operation response error`, func() {
		account := "testString"
		getRetentionPolicyPath := "/api/v1/retentions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRetentionPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRetentionPolicy with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetRetentionPolicyOptions model
				getRetentionPolicyOptionsModel := new(containerregistryv1.GetRetentionPolicyOptions)
				getRetentionPolicyOptionsModel.Namespace = core.StringPtr("testString")
				getRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRetentionPolicy(getRetentionPolicyOptions *GetRetentionPolicyOptions)`, func() {
		account := "testString"
		getRetentionPolicyPath := "/api/v1/retentions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRetentionPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"images_per_repo": 13, "namespace": "Namespace", "retain_untagged": true}`)
				}))
			})
			It(`Invoke GetRetentionPolicy successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the GetRetentionPolicyOptions model
				getRetentionPolicyOptionsModel := new(containerregistryv1.GetRetentionPolicyOptions)
				getRetentionPolicyOptionsModel.Namespace = core.StringPtr("testString")
				getRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.GetRetentionPolicyWithContext(ctx, getRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.GetRetentionPolicyWithContext(ctx, getRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRetentionPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"images_per_repo": 13, "namespace": "Namespace", "retain_untagged": true}`)
				}))
			})
			It(`Invoke GetRetentionPolicy successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.GetRetentionPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRetentionPolicyOptions model
				getRetentionPolicyOptionsModel := new(containerregistryv1.GetRetentionPolicyOptions)
				getRetentionPolicyOptionsModel.Namespace = core.StringPtr("testString")
				getRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRetentionPolicy with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetRetentionPolicyOptions model
				getRetentionPolicyOptionsModel := new(containerregistryv1.GetRetentionPolicyOptions)
				getRetentionPolicyOptionsModel.Namespace = core.StringPtr("testString")
				getRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRetentionPolicyOptions model with no property values
				getRetentionPolicyOptionsModelNew := new(containerregistryv1.GetRetentionPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetRetentionPolicy successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetRetentionPolicyOptions model
				getRetentionPolicyOptionsModel := new(containerregistryv1.GetRetentionPolicyOptions)
				getRetentionPolicyOptionsModel.Namespace = core.StringPtr("testString")
				getRetentionPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.GetRetentionPolicy(getRetentionPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		account := "testString"
		getSettingsPath := "/api/v1/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(containerregistryv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		account := "testString"
		getSettingsPath := "/api/v1/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"platform_metrics": false}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(containerregistryv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"platform_metrics": false}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(containerregistryv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(containerregistryv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(containerregistryv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
		account := "testString"
		updateSettingsPath := "/api/v1/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateSettings successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := containerRegistryService.UpdateSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(containerregistryv1.UpdateSettingsOptions)
				updateSettingsOptionsModel.PlatformMetrics = core.BoolPtr(true)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = containerRegistryService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateSettings with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(containerregistryv1.UpdateSettingsOptions)
				updateSettingsOptionsModel.PlatformMetrics = core.BoolPtr(true)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := containerRegistryService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteImageTag(deleteImageTagOptions *DeleteImageTagOptions) - Operation response error`, func() {
		account := "testString"
		deleteImageTagPath := "/api/v1/tags/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteImageTagPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteImageTag with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the DeleteImageTagOptions model
				deleteImageTagOptionsModel := new(containerregistryv1.DeleteImageTagOptions)
				deleteImageTagOptionsModel.Image = core.StringPtr("testString")
				deleteImageTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.DeleteImageTag(deleteImageTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.DeleteImageTag(deleteImageTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteImageTag(deleteImageTagOptions *DeleteImageTagOptions)`, func() {
		account := "testString"
		deleteImageTagPath := "/api/v1/tags/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteImageTagPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Untagged": "Untagged"}`)
				}))
			})
			It(`Invoke DeleteImageTag successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteImageTagOptions model
				deleteImageTagOptionsModel := new(containerregistryv1.DeleteImageTagOptions)
				deleteImageTagOptionsModel.Image = core.StringPtr("testString")
				deleteImageTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.DeleteImageTagWithContext(ctx, deleteImageTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.DeleteImageTag(deleteImageTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.DeleteImageTagWithContext(ctx, deleteImageTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteImageTagPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Untagged": "Untagged"}`)
				}))
			})
			It(`Invoke DeleteImageTag successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.DeleteImageTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteImageTagOptions model
				deleteImageTagOptionsModel := new(containerregistryv1.DeleteImageTagOptions)
				deleteImageTagOptionsModel.Image = core.StringPtr("testString")
				deleteImageTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.DeleteImageTag(deleteImageTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteImageTag with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the DeleteImageTagOptions model
				deleteImageTagOptionsModel := new(containerregistryv1.DeleteImageTagOptions)
				deleteImageTagOptionsModel.Image = core.StringPtr("testString")
				deleteImageTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.DeleteImageTag(deleteImageTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteImageTagOptions model with no property values
				deleteImageTagOptionsModelNew := new(containerregistryv1.DeleteImageTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.DeleteImageTag(deleteImageTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteImageTag successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the DeleteImageTagOptions model
				deleteImageTagOptionsModel := new(containerregistryv1.DeleteImageTagOptions)
				deleteImageTagOptionsModel.Image = core.StringPtr("testString")
				deleteImageTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.DeleteImageTag(deleteImageTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeletedImages(listDeletedImagesOptions *ListDeletedImagesOptions) - Operation response error`, func() {
		account := "testString"
		listDeletedImagesPath := "/api/v1/trash"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeletedImagesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDeletedImages with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListDeletedImagesOptions model
				listDeletedImagesOptionsModel := new(containerregistryv1.ListDeletedImagesOptions)
				listDeletedImagesOptionsModel.Namespace = core.StringPtr("testString")
				listDeletedImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.ListDeletedImages(listDeletedImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.ListDeletedImages(listDeletedImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeletedImages(listDeletedImagesOptions *ListDeletedImagesOptions)`, func() {
		account := "testString"
		listDeletedImagesPath := "/api/v1/trash"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeletedImagesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": {"daysUntilExpiry": 15, "tags": ["Tags"]}}`)
				}))
			})
			It(`Invoke ListDeletedImages successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the ListDeletedImagesOptions model
				listDeletedImagesOptionsModel := new(containerregistryv1.ListDeletedImagesOptions)
				listDeletedImagesOptionsModel.Namespace = core.StringPtr("testString")
				listDeletedImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.ListDeletedImagesWithContext(ctx, listDeletedImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.ListDeletedImages(listDeletedImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.ListDeletedImagesWithContext(ctx, listDeletedImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeletedImagesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": {"daysUntilExpiry": 15, "tags": ["Tags"]}}`)
				}))
			})
			It(`Invoke ListDeletedImages successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.ListDeletedImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDeletedImagesOptions model
				listDeletedImagesOptionsModel := new(containerregistryv1.ListDeletedImagesOptions)
				listDeletedImagesOptionsModel.Namespace = core.StringPtr("testString")
				listDeletedImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.ListDeletedImages(listDeletedImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDeletedImages with error: Operation request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListDeletedImagesOptions model
				listDeletedImagesOptionsModel := new(containerregistryv1.ListDeletedImagesOptions)
				listDeletedImagesOptionsModel.Namespace = core.StringPtr("testString")
				listDeletedImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.ListDeletedImages(listDeletedImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDeletedImages successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the ListDeletedImagesOptions model
				listDeletedImagesOptionsModel := new(containerregistryv1.ListDeletedImagesOptions)
				listDeletedImagesOptionsModel.Namespace = core.StringPtr("testString")
				listDeletedImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.ListDeletedImages(listDeletedImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RestoreTags(restoreTagsOptions *RestoreTagsOptions) - Operation response error`, func() {
		account := "testString"
		restoreTagsPath := "/api/v1/trash/testString/restoretags"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreTagsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RestoreTags with error: Operation response processing error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the RestoreTagsOptions model
				restoreTagsOptionsModel := new(containerregistryv1.RestoreTagsOptions)
				restoreTagsOptionsModel.Digest = core.StringPtr("testString")
				restoreTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := containerRegistryService.RestoreTags(restoreTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				containerRegistryService.EnableRetries(0, 0)
				result, response, operationErr = containerRegistryService.RestoreTags(restoreTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RestoreTags(restoreTagsOptions *RestoreTagsOptions)`, func() {
		account := "testString"
		restoreTagsPath := "/api/v1/trash/testString/restoretags"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreTagsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"successful": ["Successful"], "unsuccessful": ["Unsuccessful"]}`)
				}))
			})
			It(`Invoke RestoreTags successfully with retries`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())
				containerRegistryService.EnableRetries(0, 0)

				// Construct an instance of the RestoreTagsOptions model
				restoreTagsOptionsModel := new(containerregistryv1.RestoreTagsOptions)
				restoreTagsOptionsModel.Digest = core.StringPtr("testString")
				restoreTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := containerRegistryService.RestoreTagsWithContext(ctx, restoreTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				containerRegistryService.DisableRetries()
				result, response, operationErr := containerRegistryService.RestoreTags(restoreTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = containerRegistryService.RestoreTagsWithContext(ctx, restoreTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreTagsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"successful": ["Successful"], "unsuccessful": ["Unsuccessful"]}`)
				}))
			})
			It(`Invoke RestoreTags successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := containerRegistryService.RestoreTags(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RestoreTagsOptions model
				restoreTagsOptionsModel := new(containerregistryv1.RestoreTagsOptions)
				restoreTagsOptionsModel.Digest = core.StringPtr("testString")
				restoreTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = containerRegistryService.RestoreTags(restoreTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RestoreTags with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the RestoreTagsOptions model
				restoreTagsOptionsModel := new(containerregistryv1.RestoreTagsOptions)
				restoreTagsOptionsModel.Digest = core.StringPtr("testString")
				restoreTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := containerRegistryService.RestoreTags(restoreTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RestoreTagsOptions model with no property values
				restoreTagsOptionsModelNew := new(containerregistryv1.RestoreTagsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = containerRegistryService.RestoreTags(restoreTagsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke RestoreTags successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the RestoreTagsOptions model
				restoreTagsOptionsModel := new(containerregistryv1.RestoreTagsOptions)
				restoreTagsOptionsModel.Digest = core.StringPtr("testString")
				restoreTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := containerRegistryService.RestoreTags(restoreTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RestoreImage(restoreImageOptions *RestoreImageOptions)`, func() {
		account := "testString"
		restoreImagePath := "/api/v1/trash/testString/restore"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreImagePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Account"]).ToNot(BeNil())
					Expect(req.Header["Account"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke RestoreImage successfully`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := containerRegistryService.RestoreImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RestoreImageOptions model
				restoreImageOptionsModel := new(containerregistryv1.RestoreImageOptions)
				restoreImageOptionsModel.Image = core.StringPtr("testString")
				restoreImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = containerRegistryService.RestoreImage(restoreImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RestoreImage with error: Operation validation and request error`, func() {
				containerRegistryService, serviceErr := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Account: core.StringPtr(account),
				})
				Expect(serviceErr).To(BeNil())
				Expect(containerRegistryService).ToNot(BeNil())

				// Construct an instance of the RestoreImageOptions model
				restoreImageOptionsModel := new(containerregistryv1.RestoreImageOptions)
				restoreImageOptionsModel.Image = core.StringPtr("testString")
				restoreImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := containerRegistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := containerRegistryService.RestoreImage(restoreImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RestoreImageOptions model with no property values
				restoreImageOptionsModelNew := new(containerregistryv1.RestoreImageOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = containerRegistryService.RestoreImage(restoreImageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			account := "testString"
			containerRegistryService, _ := containerregistryv1.NewContainerRegistryV1(&containerregistryv1.ContainerRegistryV1Options{
				URL:           "http://containerregistryv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Account: core.StringPtr(account),
			})
			It(`Invoke NewAnalyzeRetentionPolicyOptions successfully`, func() {
				// Construct an instance of the AnalyzeRetentionPolicyOptions model
				analyzeRetentionPolicyOptionsNamespace := "birds"
				analyzeRetentionPolicyOptionsModel := containerRegistryService.NewAnalyzeRetentionPolicyOptions(analyzeRetentionPolicyOptionsNamespace)
				analyzeRetentionPolicyOptionsModel.SetNamespace("birds")
				analyzeRetentionPolicyOptionsModel.SetImagesPerRepo(int64(10))
				analyzeRetentionPolicyOptionsModel.SetRetainUntagged(false)
				analyzeRetentionPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(analyzeRetentionPolicyOptionsModel).ToNot(BeNil())
				Expect(analyzeRetentionPolicyOptionsModel.Namespace).To(Equal(core.StringPtr("birds")))
				Expect(analyzeRetentionPolicyOptionsModel.ImagesPerRepo).To(Equal(core.Int64Ptr(int64(10))))
				Expect(analyzeRetentionPolicyOptionsModel.RetainUntagged).To(Equal(core.BoolPtr(false)))
				Expect(analyzeRetentionPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAssignNamespaceOptions successfully`, func() {
				// Construct an instance of the AssignNamespaceOptions model
				xAuthResourceGroup := "testString"
				name := "testString"
				assignNamespaceOptionsModel := containerRegistryService.NewAssignNamespaceOptions(xAuthResourceGroup, name)
				assignNamespaceOptionsModel.SetXAuthResourceGroup("testString")
				assignNamespaceOptionsModel.SetName("testString")
				assignNamespaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(assignNamespaceOptionsModel).ToNot(BeNil())
				Expect(assignNamespaceOptionsModel.XAuthResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(assignNamespaceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(assignNamespaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewBulkDeleteImagesOptions successfully`, func() {
				// Construct an instance of the BulkDeleteImagesOptions model
				bulkDelete := []string{"us.icr.io/birds/woodpecker@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4bbbb", "us.icr.io/birds/bird@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4dddd"}
				bulkDeleteImagesOptionsModel := containerRegistryService.NewBulkDeleteImagesOptions(bulkDelete)
				bulkDeleteImagesOptionsModel.SetBulkDelete([]string{"us.icr.io/birds/woodpecker@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4bbbb", "us.icr.io/birds/bird@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4dddd"})
				bulkDeleteImagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(bulkDeleteImagesOptionsModel).ToNot(BeNil())
				Expect(bulkDeleteImagesOptionsModel.BulkDelete).To(Equal([]string{"us.icr.io/birds/woodpecker@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4bbbb", "us.icr.io/birds/bird@sha256:38f97dd92769b18ca82ad9ab6667af47306e66fea5b446937eea68b10ab4dddd"}))
				Expect(bulkDeleteImagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateNamespaceOptions successfully`, func() {
				// Construct an instance of the CreateNamespaceOptions model
				name := "testString"
				createNamespaceOptionsModel := containerRegistryService.NewCreateNamespaceOptions(name)
				createNamespaceOptionsModel.SetName("testString")
				createNamespaceOptionsModel.SetXAuthResourceGroup("testString")
				createNamespaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createNamespaceOptionsModel).ToNot(BeNil())
				Expect(createNamespaceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createNamespaceOptionsModel.XAuthResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(createNamespaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteImageOptions successfully`, func() {
				// Construct an instance of the DeleteImageOptions model
				image := "testString"
				deleteImageOptionsModel := containerRegistryService.NewDeleteImageOptions(image)
				deleteImageOptionsModel.SetImage("testString")
				deleteImageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteImageOptionsModel).ToNot(BeNil())
				Expect(deleteImageOptionsModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(deleteImageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteImageTagOptions successfully`, func() {
				// Construct an instance of the DeleteImageTagOptions model
				image := "testString"
				deleteImageTagOptionsModel := containerRegistryService.NewDeleteImageTagOptions(image)
				deleteImageTagOptionsModel.SetImage("testString")
				deleteImageTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteImageTagOptionsModel).ToNot(BeNil())
				Expect(deleteImageTagOptionsModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(deleteImageTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteNamespaceOptions successfully`, func() {
				// Construct an instance of the DeleteNamespaceOptions model
				name := "testString"
				deleteNamespaceOptionsModel := containerRegistryService.NewDeleteNamespaceOptions(name)
				deleteNamespaceOptionsModel.SetName("testString")
				deleteNamespaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteNamespaceOptionsModel).ToNot(BeNil())
				Expect(deleteNamespaceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(deleteNamespaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAuthOptions successfully`, func() {
				// Construct an instance of the GetAuthOptions model
				getAuthOptionsModel := containerRegistryService.NewGetAuthOptions()
				getAuthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAuthOptionsModel).ToNot(BeNil())
				Expect(getAuthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetImageManifestOptions successfully`, func() {
				// Construct an instance of the GetImageManifestOptions model
				image := "testString"
				getImageManifestOptionsModel := containerRegistryService.NewGetImageManifestOptions(image)
				getImageManifestOptionsModel.SetImage("testString")
				getImageManifestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getImageManifestOptionsModel).ToNot(BeNil())
				Expect(getImageManifestOptionsModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(getImageManifestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMessagesOptions successfully`, func() {
				// Construct an instance of the GetMessagesOptions model
				getMessagesOptionsModel := containerRegistryService.NewGetMessagesOptions()
				getMessagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMessagesOptionsModel).ToNot(BeNil())
				Expect(getMessagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPlansOptions successfully`, func() {
				// Construct an instance of the GetPlansOptions model
				getPlansOptionsModel := containerRegistryService.NewGetPlansOptions()
				getPlansOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPlansOptionsModel).ToNot(BeNil())
				Expect(getPlansOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetQuotaOptions successfully`, func() {
				// Construct an instance of the GetQuotaOptions model
				getQuotaOptionsModel := containerRegistryService.NewGetQuotaOptions()
				getQuotaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getQuotaOptionsModel).ToNot(BeNil())
				Expect(getQuotaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRetentionPolicyOptions successfully`, func() {
				// Construct an instance of the GetRetentionPolicyOptions model
				namespace := "testString"
				getRetentionPolicyOptionsModel := containerRegistryService.NewGetRetentionPolicyOptions(namespace)
				getRetentionPolicyOptionsModel.SetNamespace("testString")
				getRetentionPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRetentionPolicyOptionsModel).ToNot(BeNil())
				Expect(getRetentionPolicyOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(getRetentionPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := containerRegistryService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInspectImageOptions successfully`, func() {
				// Construct an instance of the InspectImageOptions model
				image := "testString"
				inspectImageOptionsModel := containerRegistryService.NewInspectImageOptions(image)
				inspectImageOptionsModel.SetImage("testString")
				inspectImageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(inspectImageOptionsModel).ToNot(BeNil())
				Expect(inspectImageOptionsModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(inspectImageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDeletedImagesOptions successfully`, func() {
				// Construct an instance of the ListDeletedImagesOptions model
				listDeletedImagesOptionsModel := containerRegistryService.NewListDeletedImagesOptions()
				listDeletedImagesOptionsModel.SetNamespace("testString")
				listDeletedImagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDeletedImagesOptionsModel).ToNot(BeNil())
				Expect(listDeletedImagesOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(listDeletedImagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListImageDigestsOptions successfully`, func() {
				// Construct an instance of the ListImageDigestsOptions model
				listImageDigestsOptionsModel := containerRegistryService.NewListImageDigestsOptions()
				listImageDigestsOptionsModel.SetExcludeTagged(false)
				listImageDigestsOptionsModel.SetExcludeVa(false)
				listImageDigestsOptionsModel.SetIncludeIBM(false)
				listImageDigestsOptionsModel.SetRepositories([]string{"testString"})
				listImageDigestsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listImageDigestsOptionsModel).ToNot(BeNil())
				Expect(listImageDigestsOptionsModel.ExcludeTagged).To(Equal(core.BoolPtr(false)))
				Expect(listImageDigestsOptionsModel.ExcludeVa).To(Equal(core.BoolPtr(false)))
				Expect(listImageDigestsOptionsModel.IncludeIBM).To(Equal(core.BoolPtr(false)))
				Expect(listImageDigestsOptionsModel.Repositories).To(Equal([]string{"testString"}))
				Expect(listImageDigestsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListImagesOptions successfully`, func() {
				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := containerRegistryService.NewListImagesOptions()
				listImagesOptionsModel.SetNamespace("testString")
				listImagesOptionsModel.SetIncludeIBM(true)
				listImagesOptionsModel.SetIncludePrivate(true)
				listImagesOptionsModel.SetIncludeManifestLists(true)
				listImagesOptionsModel.SetVulnerabilities(true)
				listImagesOptionsModel.SetRepository("testString")
				listImagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listImagesOptionsModel).ToNot(BeNil())
				Expect(listImagesOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(listImagesOptionsModel.IncludeIBM).To(Equal(core.BoolPtr(true)))
				Expect(listImagesOptionsModel.IncludePrivate).To(Equal(core.BoolPtr(true)))
				Expect(listImagesOptionsModel.IncludeManifestLists).To(Equal(core.BoolPtr(true)))
				Expect(listImagesOptionsModel.Vulnerabilities).To(Equal(core.BoolPtr(true)))
				Expect(listImagesOptionsModel.Repository).To(Equal(core.StringPtr("testString")))
				Expect(listImagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListNamespaceDetailsOptions successfully`, func() {
				// Construct an instance of the ListNamespaceDetailsOptions model
				listNamespaceDetailsOptionsModel := containerRegistryService.NewListNamespaceDetailsOptions()
				listNamespaceDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNamespaceDetailsOptionsModel).ToNot(BeNil())
				Expect(listNamespaceDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListNamespacesOptions successfully`, func() {
				// Construct an instance of the ListNamespacesOptions model
				listNamespacesOptionsModel := containerRegistryService.NewListNamespacesOptions()
				listNamespacesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNamespacesOptionsModel).ToNot(BeNil())
				Expect(listNamespacesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRetentionPoliciesOptions successfully`, func() {
				// Construct an instance of the ListRetentionPoliciesOptions model
				listRetentionPoliciesOptionsModel := containerRegistryService.NewListRetentionPoliciesOptions()
				listRetentionPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRetentionPoliciesOptionsModel).ToNot(BeNil())
				Expect(listRetentionPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRestoreImageOptions successfully`, func() {
				// Construct an instance of the RestoreImageOptions model
				image := "testString"
				restoreImageOptionsModel := containerRegistryService.NewRestoreImageOptions(image)
				restoreImageOptionsModel.SetImage("testString")
				restoreImageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreImageOptionsModel).ToNot(BeNil())
				Expect(restoreImageOptionsModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(restoreImageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRestoreTagsOptions successfully`, func() {
				// Construct an instance of the RestoreTagsOptions model
				digest := "testString"
				restoreTagsOptionsModel := containerRegistryService.NewRestoreTagsOptions(digest)
				restoreTagsOptionsModel.SetDigest("testString")
				restoreTagsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreTagsOptionsModel).ToNot(BeNil())
				Expect(restoreTagsOptionsModel.Digest).To(Equal(core.StringPtr("testString")))
				Expect(restoreTagsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRetentionPolicy successfully`, func() {
				namespace := "testString"
				_model, err := containerRegistryService.NewRetentionPolicy(namespace)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSetRetentionPolicyOptions successfully`, func() {
				// Construct an instance of the SetRetentionPolicyOptions model
				setRetentionPolicyOptionsNamespace := "birds"
				setRetentionPolicyOptionsModel := containerRegistryService.NewSetRetentionPolicyOptions(setRetentionPolicyOptionsNamespace)
				setRetentionPolicyOptionsModel.SetNamespace("birds")
				setRetentionPolicyOptionsModel.SetImagesPerRepo(int64(10))
				setRetentionPolicyOptionsModel.SetRetainUntagged(false)
				setRetentionPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setRetentionPolicyOptionsModel).ToNot(BeNil())
				Expect(setRetentionPolicyOptionsModel.Namespace).To(Equal(core.StringPtr("birds")))
				Expect(setRetentionPolicyOptionsModel.ImagesPerRepo).To(Equal(core.Int64Ptr(int64(10))))
				Expect(setRetentionPolicyOptionsModel.RetainUntagged).To(Equal(core.BoolPtr(false)))
				Expect(setRetentionPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTagImageOptions successfully`, func() {
				// Construct an instance of the TagImageOptions model
				fromimage := "testString"
				toimage := "testString"
				tagImageOptionsModel := containerRegistryService.NewTagImageOptions(fromimage, toimage)
				tagImageOptionsModel.SetFromimage("testString")
				tagImageOptionsModel.SetToimage("testString")
				tagImageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(tagImageOptionsModel).ToNot(BeNil())
				Expect(tagImageOptionsModel.Fromimage).To(Equal(core.StringPtr("testString")))
				Expect(tagImageOptionsModel.Toimage).To(Equal(core.StringPtr("testString")))
				Expect(tagImageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAuthOptions successfully`, func() {
				// Construct an instance of the UpdateAuthOptions model
				updateAuthOptionsModel := containerRegistryService.NewUpdateAuthOptions()
				updateAuthOptionsModel.SetIamAuthz(true)
				updateAuthOptionsModel.SetPrivateOnly(true)
				updateAuthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAuthOptionsModel).ToNot(BeNil())
				Expect(updateAuthOptionsModel.IamAuthz).To(Equal(core.BoolPtr(true)))
				Expect(updateAuthOptionsModel.PrivateOnly).To(Equal(core.BoolPtr(true)))
				Expect(updateAuthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePlansOptions successfully`, func() {
				// Construct an instance of the UpdatePlansOptions model
				updatePlansOptionsModel := containerRegistryService.NewUpdatePlansOptions()
				updatePlansOptionsModel.SetPlan("Standard")
				updatePlansOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePlansOptionsModel).ToNot(BeNil())
				Expect(updatePlansOptionsModel.Plan).To(Equal(core.StringPtr("Standard")))
				Expect(updatePlansOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateQuotaOptions successfully`, func() {
				// Construct an instance of the UpdateQuotaOptions model
				updateQuotaOptionsModel := containerRegistryService.NewUpdateQuotaOptions()
				updateQuotaOptionsModel.SetStorageMegabytes(int64(26))
				updateQuotaOptionsModel.SetTrafficMegabytes(int64(480))
				updateQuotaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateQuotaOptionsModel).ToNot(BeNil())
				Expect(updateQuotaOptionsModel.StorageMegabytes).To(Equal(core.Int64Ptr(int64(26))))
				Expect(updateQuotaOptionsModel.TrafficMegabytes).To(Equal(core.Int64Ptr(int64(480))))
				Expect(updateQuotaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSettingsOptions successfully`, func() {
				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := containerRegistryService.NewUpdateSettingsOptions()
				updateSettingsOptionsModel.SetPlatformMetrics(true)
				updateSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSettingsOptionsModel).ToNot(BeNil())
				Expect(updateSettingsOptionsModel.PlatformMetrics).To(Equal(core.BoolPtr(true)))
				Expect(updateSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
