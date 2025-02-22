/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    BatchHintsResponseBody,
    BatchHintsResponseBodyFromJSON,
    BatchHintsResponseBodyToJSON,
    BatchStatusResponseBody,
    BatchStatusResponseBodyFromJSON,
    BatchStatusResponseBodyToJSON,
    BatchSubmitNotAvailableResponseBody,
    BatchSubmitNotAvailableResponseBodyFromJSON,
    BatchSubmitNotAvailableResponseBodyToJSON,
    BatchSubmitNotValidResponseBody,
    BatchSubmitNotValidResponseBodyFromJSON,
    BatchSubmitNotValidResponseBodyToJSON,
    BatchSubmitRequestBody,
    BatchSubmitRequestBodyFromJSON,
    BatchSubmitRequestBodyToJSON,
    BatchSubmitResponseBody,
    BatchSubmitResponseBodyFromJSON,
    BatchSubmitResponseBodyToJSON,
} from '../models';

export interface BatchSubmitRequest {
    submitRequestBody: BatchSubmitRequestBody;
}

/**
 * no description
 */
export class BatchApi extends runtime.BaseAPI {

    /**
     * Retrieve form hints
     * hints batch
     */
    async batchHintsRaw(): Promise<runtime.ApiResponse<BatchHintsResponseBody>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/batch/hints`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => BatchHintsResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve form hints
     * hints batch
     */
    async batchHints(): Promise<BatchHintsResponseBody> {
        const response = await this.batchHintsRaw();
        return await response.value();
    }

    /**
     * Retrieve status of current batch operation.
     * status batch
     */
    async batchStatusRaw(): Promise<runtime.ApiResponse<BatchStatusResponseBody>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/batch`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => BatchStatusResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve status of current batch operation.
     * status batch
     */
    async batchStatus(): Promise<BatchStatusResponseBody> {
        const response = await this.batchStatusRaw();
        return await response.value();
    }

    /**
     * Submit a new batch
     * submit batch
     */
    async batchSubmitRaw(requestParameters: BatchSubmitRequest): Promise<runtime.ApiResponse<BatchSubmitResponseBody>> {
        if (requestParameters.submitRequestBody === null || requestParameters.submitRequestBody === undefined) {
            throw new runtime.RequiredError('submitRequestBody','Required parameter requestParameters.submitRequestBody was null or undefined when calling batchSubmit.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/batch`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: BatchSubmitRequestBodyToJSON(requestParameters.submitRequestBody),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => BatchSubmitResponseBodyFromJSON(jsonValue));
    }

    /**
     * Submit a new batch
     * submit batch
     */
    async batchSubmit(requestParameters: BatchSubmitRequest): Promise<BatchSubmitResponseBody> {
        const response = await this.batchSubmitRaw(requestParameters);
        return await response.value();
    }

}
