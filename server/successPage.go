package server

var successPage []byte = []byte(`<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <title>dexctl</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
		* {
		  box-sizing: border-box;
		}

		body {
		  margin: 0;
		}

		.dex-container {
		  color: #333;
		  margin: 45px auto;
		  max-width: 500px;
		  min-width: 320px;
		  text-align: center;
		}

		.dex-btn {
		  border-radius: 4px;
		  border: 0;
		  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.25), 0 0 1px rgba(0, 0, 0, 0.25);
		  cursor: pointer;
		  font-size: 16px;
		  padding: 0;
		}

		.dex-btn:focus {
		  outline: none;
		}

		.dex-btn:active {
		  box-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.125);
		  outline: none;
		}

		.dex-btn-icon {
		  background-position: center;
		  background-repeat: no-repeat;
		  background-size: 24px;
		  border-radius: 4px 0 0 4px;
		  float: left;
		  height: 36px;
		  margin-right: 5px;
		  width: 36px;
		}

		.dex-btn-icon--google {
		  background-color: #FFFFFF;
		  background-image: url(../static/img/google-icon.svg);;
		}

		.dex-btn-icon--local {
		  background-color: #84B6EF;
		  background-image: url(../static/img/email-icon.svg);
		}

		.dex-btn-icon--gitea {
		  background-color: #F5F5F5;
		  background-image: url(../static/img/gitea-icon.svg);
		}

		.dex-btn-icon--github {
		  background-color: #F5F5F5;
		  background-image: url(../static/img/github-icon.svg);
		}

		.dex-btn-icon--gitlab {
		  background-color: #F5F5F5;
		  background-image: url(../static/img/gitlab-icon.svg);
		  background-size: contain;
		}

		.dex-btn-icon--keystone {
		  background-color: #F5F5F5;
		  background-image: url(../static/img/keystone-icon.svg);
		  background-size: contain;
		}

		.dex-btn-icon--oidc {
		  background-color: #EBEBEE;
		  background-image: url(../static/img/oidc-icon.svg);
		  background-size: contain;
		}

		.dex-btn-icon--bitbucket-cloud {
		  background-color: #205081;
		  background-image: url(../static/img/bitbucket-icon.svg);
		}

		.dex-btn-icon--atlassian-crowd {
		  background-color: #CFDCEA;
		  background-image: url(../static/img/atlassian-crowd-icon.svg);
		}

		.dex-btn-icon--ldap {
		  background-color: #84B6EF;
		  background-image: url(../static/img/ldap-icon.svg);
		}

		.dex-btn-icon--saml {
		  background-color: #84B6EF;
		  background-image: url(../static/img/saml-icon.svg);
		}

		.dex-btn-icon--linkedin {
		  background-image: url(../static/img/linkedin-icon.svg);
		  background-size: contain;
		}

		.dex-btn-icon--microsoft {
		  background-image: url(../static/img/microsoft-icon.svg);
		}

		.dex-btn-text {
		  font-weight: 600;
		  line-height: 36px;
		  padding: 6px 12px;
		  text-align: center;
		}

		.dex-subtle-text {
		  color: #999;
		  font-size: 12px;
		}

		.dex-separator {
		  color: #999;
		}

		.dex-list {
		  color: #999;
		  display: inline-block;
		  font-size: 12px;
		  list-style: circle;
		  text-align: left;
		}

		.dex-error-box {
		  background-color: #DD1327;
		  color: #fff;
		  font-size: 14px;
		  font-weight: normal;
		  max-width: 320px;
		  padding: 4px 0;
		}

		.dex-error-box {
		  margin: 20px auto;
		}

		.theme-body {
		  background-color: #efefef;
		  color: #333;
		  font-family: 'Source Sans Pro', Helvetica, sans-serif;
		}

		.theme-navbar {
		  background-color: #fff;
		  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.2);
		  color: #333;
		  font-size: 13px;
		  font-weight: 100;
		  height: 46px;
		  overflow: hidden;
		  padding: 0 10px;
		}

		.theme-navbar__logo-wrap {
		  display: inline-block;
		  height: 100%;
		  overflow: hidden;
		  padding: 10px 15px;
		  width: 300px;
		}

		.theme-navbar__logo {
		  height: 100%;
		  max-height: 25px;
		}

		.theme-heading {
		  font-size: 20px;
		  font-weight: 500;
		  margin-bottom: 10px;
		  margin-top: 0;
		}

		.theme-panel {
		  background-color: #fff;
		  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.5);
		  padding: 30px;
		}

		.theme-btn-provider {
		  background-color: #fff;
		  color: #333;
		  min-width: 250px;
		}

		.theme-btn-provider:hover {
		  color: #999;
		}

		.theme-btn--primary {
		  background-color: #333;
		  border: none;
		  color: #fff;
		  min-width: 200px;
		  padding: 6px 12px;
		}

		.theme-btn--primary:hover {
		  background-color: #666;
		  color: #fff;
		}

		.theme-btn--success {
		  background-color: #2FC98E;
		  color: #fff;
		  width: 250px;
		}

		.theme-btn--success:hover {
		  background-color: #49E3A8;
		}

		.theme-form-row {
		  display: block;
		  margin: 20px auto;
		}

		.theme-form-input {
		  border-radius: 4px;
		  border: 1px solid #CCC;
		  box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
		  color: #666;
		  display: block;
		  font-size: 14px;
		  height: 36px;
		  line-height: 1.42857143;
		  margin: auto;
		  padding: 6px 12px;
		  width: 250px;
		}

		.theme-form-input:focus,
		.theme-form-input:active {
		  border-color: #66AFE9;
		  outline: none;
		}

		.theme-form-label {
		  font-size: 13px;
		  font-weight: 600;
		  margin: 4px auto;
		  position: relative;
		  text-align: left;
		  width: 250px;
		}

		.theme-link-back {
		  margin-top: 4px;
		}

    </style>
  </head>

  <body class="theme-body">
    <div class="theme-navbar">
    </div>

    <div class="dex-container">


<div class="theme-panel">
  <h2 class="theme-heading">Success!</h2>
  <p>go back to dexctl</p>
</div>

    </div>
  </body>
</html>`)
