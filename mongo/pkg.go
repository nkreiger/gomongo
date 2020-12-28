// package mongo handles the creation and execution of mongo queries
// the actual results are handled by the reporter, therefore,
// queries can import reporter to return the appropriate mapped structs, but
// reporter cannot access queries
//
// currently handles mongo queries, specifically, of type find and findOne
// see mongo documentation for more information
package mongo