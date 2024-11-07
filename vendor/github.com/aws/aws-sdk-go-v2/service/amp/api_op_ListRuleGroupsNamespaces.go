// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of rule groups namespaces in a workspace.
func (c *Client) ListRuleGroupsNamespaces(ctx context.Context, params *ListRuleGroupsNamespacesInput, optFns ...func(*Options)) (*ListRuleGroupsNamespacesOutput, error) {
	if params == nil {
		params = &ListRuleGroupsNamespacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRuleGroupsNamespaces", params, optFns, c.addOperationListRuleGroupsNamespacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRuleGroupsNamespacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListRuleGroupsNamespaces operation.
type ListRuleGroupsNamespacesInput struct {

	// The ID of the workspace containing the rule groups namespaces.
	//
	// This member is required.
	WorkspaceId *string

	// The maximum number of results to return. The default is 100.
	MaxResults *int32

	// Use this parameter to filter the rule groups namespaces that are returned. Only
	// the namespaces with names that begin with the value that you specify are
	// returned.
	Name *string

	// The token for the next set of items to return. You receive this token from a
	// previous call, and use it to get the next page of results. The other parameters
	// must be the same as the initial call.
	//
	// For example, if your initial request has maxResults of 10, and there are 12
	// rule groups namespaces to return, then your initial request will return 10 and a
	// nextToken . Using the next token in a subsequent call will return the remaining
	// 2 namespaces.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the output of a ListRuleGroupsNamespaces operation.
type ListRuleGroupsNamespacesOutput struct {

	// The returned list of rule groups namespaces.
	//
	// This member is required.
	RuleGroupsNamespaces []types.RuleGroupsNamespaceSummary

	// A token indicating that there are more results to retrieve. You can use this
	// token as part of your next ListRuleGroupsNamespaces request to retrieve those
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRuleGroupsNamespacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRuleGroupsNamespaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRuleGroupsNamespaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRuleGroupsNamespaces"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpListRuleGroupsNamespacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRuleGroupsNamespaces(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// ListRuleGroupsNamespacesPaginatorOptions is the paginator options for
// ListRuleGroupsNamespaces
type ListRuleGroupsNamespacesPaginatorOptions struct {
	// The maximum number of results to return. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRuleGroupsNamespacesPaginator is a paginator for ListRuleGroupsNamespaces
type ListRuleGroupsNamespacesPaginator struct {
	options   ListRuleGroupsNamespacesPaginatorOptions
	client    ListRuleGroupsNamespacesAPIClient
	params    *ListRuleGroupsNamespacesInput
	nextToken *string
	firstPage bool
}

// NewListRuleGroupsNamespacesPaginator returns a new
// ListRuleGroupsNamespacesPaginator
func NewListRuleGroupsNamespacesPaginator(client ListRuleGroupsNamespacesAPIClient, params *ListRuleGroupsNamespacesInput, optFns ...func(*ListRuleGroupsNamespacesPaginatorOptions)) *ListRuleGroupsNamespacesPaginator {
	if params == nil {
		params = &ListRuleGroupsNamespacesInput{}
	}

	options := ListRuleGroupsNamespacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRuleGroupsNamespacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRuleGroupsNamespacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRuleGroupsNamespaces page.
func (p *ListRuleGroupsNamespacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRuleGroupsNamespacesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListRuleGroupsNamespaces(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListRuleGroupsNamespacesAPIClient is a client that implements the
// ListRuleGroupsNamespaces operation.
type ListRuleGroupsNamespacesAPIClient interface {
	ListRuleGroupsNamespaces(context.Context, *ListRuleGroupsNamespacesInput, ...func(*Options)) (*ListRuleGroupsNamespacesOutput, error)
}

var _ ListRuleGroupsNamespacesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRuleGroupsNamespaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRuleGroupsNamespaces",
	}
}
