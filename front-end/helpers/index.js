(function (){
  'use strict';

  var request = require("request");
  var helpers = {};

  helpers.sessionMiddleware = function(err, req, res, next) {
    if(!req.cookies.logged_in) {
      res.session.customerId = null;
    }
  };


  helpers.errorHandler = function(err, req, res, next) {
  var ret = {
    message: err.message,
    error:   err
  };
  res.
    status(err.status || 500).
    send(ret);
};


  /* Responds with the given statusCode */
  helpers.respondStatus = function(res, statusCode) {
    res.writeHeader(statusCode);
    res.end();
  }


  /* Responds with the given body and status 200 OK  */
  helpers.respondSuccessBody = function(res, body) {
    helpers.respondStatusBody(res, 200, body);
  }

  /* Public: responds with the given body and status
   *
   * res        - response object to use as output
   * statusCode - the HTTP status code to set to the response
   * body       - (string) the body to yield to the response
   */
  helpers.respondStatusBody = function(res, statusCode, body) {
    res.writeHeader(statusCode);
    res.write(body);
    res.end();
  }


  /* Rewrites and redirects any url that doesn't end with a slash. */
  helpers.rewriteSlash = function(req, res, next) {
   if(req.url.substr(-1) == '/' && req.url.length > 1)
       res.redirect(301, req.url.slice(0, -1));
   else
       next();
  }

  /* TODO: Add documentation */
  helpers.getCustomerId = function(req, env) {
    // Check if logged in. Get customer Id
    var logged_in = req.cookies.logged_in;

    // TODO REMOVE THIS, SECURITY RISK
    if (env == "development" && req.query.custId != null) {
      return req.query.custId;
    }

    if (!logged_in) {
      if (!req.session.id) {
        throw new Error("User not logged in.");
      }
      // Use Session ID instead
      return req.session.id;
    }

    return req.session.customerId;
  }


  helpers.simpleHttpRequest = function(url, res, next) {
    request.get(url, function(error, response, body) {
      if (error) return next(error);
      helpers.respondSuccessBody(res, body);
    }.bind({res: res}));
  }
  module.exports = helpers;
}());
