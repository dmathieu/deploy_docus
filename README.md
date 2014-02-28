# Deploy Docus [![Build Status](https://travis-ci.org/dmathieu/deploy_docus.png?branch=master)](https://travis-ci.org/dmathieu/deploy_docus)

Use [GitHub's Deployments API](http://developer.github.com/v3/repos/deployments/) to automate the deployment of your app on Heroku.

**This is still in an early beta and there are bugs. Please [let us know](https://github.com/dmathieu/deploy_docus/issues).**

## How it works

A few months ago, GitHub unveiled their [Deployments API](http://developer.github.com/v3/repos/deployments/).  
This API allows us to create new deployments for a repository.  
Whenever a new deployment is created, an event is propagated. Any application can listen to that event using the [Webhooks](https://developer.github.com/webhooks/).

Deploy Docus is a simple Go application which will listen to those events, and deploy your application on Heroku whenever it receives one.  
That way, you don't have to worry about deploying your application from the command line and can start doing [ChatOps](https://speakerdeck.com/jnewland/chatops-at-github).

## Installing

Deploy Docus is meant to be run on Heroku.

Start by cloning the repository.

    git clone https://github.com/dmathieu/deploy_docus.git
    cd deploy_docus

Then, create a new Heroku application:

    heroku create --stack cedar

You will also need to create a new [github oauth application](https://github.com/settings/applications),
to allow access only to you or the members of your organization.

Set the following configuration in your Heroku application:

    BUILDPACK_URL:                  https://github.com/kr/heroku-buildpack-go.git#go1.2
    GITHUB_OAUTH_ALLOWED_ID:        <Your GitHub ID, or your organization's>
    GITHUB_OAUTH_KEY:               <The GitHub application's oauth key>
    GITHUB_OAUTH_REDIRECT_URI:      http://<your heroku app name>.herokuapp.com/oauth2callback
    GITHUB_OAUTH_SECRET:            <The GitHub application's oauth secret>
    SECRET_SESSION_TOKEN:           <A random string>

Your application will also need a PostgreSQL database.

    heroku addons:add heroku-postgresql

You can then push the application to Heroku and head to it's URL.

    git push heroku master:master

## Testing your first deployment

Once you have deployment the application, head to the `/repositories` URL inside it.  
This will ask you to login with OAuth to GitHub and show you a simple interface allowing you to create new repositories.

Click on "Create New Repository". There, a form will ask you for two fields.

* The Origin is the SSH URL to your GitHub repository. It will be pulled from there every deploy.
* The Destination is the SSH URL to your Heroku repository. It will be pushed there every deploy.

Once you create the repository, it will be added to the application's database and you will see it in your browser.  
Click on that repository's title to see it's details.

Two things matter to us in that page.

* The SSH Public key. Every repository has it's own SSH key generated automatically by the application.  
You need to add that SSH key to [a GitHub account](https://help.github.com/articles/generating-ssh-keys#step-3-add-your-ssh-key-to-github).

* The Callback URL. That URL (with the ?token=xx in the URL) is the URL to which the application expects GitHub to make the event HTTP call.  
[Add it as a webhook in your GitHub repository](https://help.github.com/articles/creating-webhooks).

You can now use something like [hubot-deploy](https://github.com/atmos/hubot-deploy) to create new deployments of your application.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## License

Deploy Docus is released under the MIT license. See [LICENSE](LICENSE)
