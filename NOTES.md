# Pre-bootstrap

## Get credentials for pushing images to Quay.io

 Visit the settings page for your Quay.io account "https://quay.io/user/<USERNAME>?tab=settings"

 You'll be prompted to authenticate, and then you'll get a screen that allows
you download credential, pick the "Kubernetes Secret" on the left hand of the
screen.

 On this screen, there is a link below "Step 1", to download your secret
"Download <USERNAME>-secret.yml", download this file, and rename it to
`~/quayio-secret.yaml`.

 Edit `~/demo-quayio-secret.yaml` and change

 ```
 metadata:
   name: <USERNAME>-pull-secret
 ```

to

 ```
 metadata:
  name: pull-secret
 ```

## Get configuration for pulling images from Quay.io

On the same screen, as the Quay.io pull credentials, click on the "Docker Configuration", and click to download the `<USERNAME>-auth.json` file.

Rename this file to `~/demo-auth.json`

## Configure the repositories that you want to watch to trigger builds.

In [the eventlistener](./tekton/eventlisteners/cicd-event-listener.yaml there
are two repositories that are used, one for the main source (which triggers the
ci pipeline in the dev environment) and a second for the `stage-environment`.

Modify the repos `bigkevmcd/taxi` and `bigkevmcd/taxi-stage-config` as
appropriate.

e.g.

 ```
 - name: dev-ci-build-from-pr
   interceptor:
     header:
     - name: Pullrequest-Action
       value: opened
     - name:  Pullrequest-Repo
       value: bigkevmcd/taxi
     objectRef:
       kind: Service
       name: demo-interceptor
       apiVersion: v1
       namespace: cicd-environment
 ```

This assumes that `bigkevmcd/taxi` is the main repository, you can configure
this as appropriate.

## Login to your OpenShift cluster

 ```
 oc login --token=<TOKEN FROM YOUR CLUSTER> --server=https://<CLUSTER ADDRESS>
 ```

## Execute the bootstrap command

 ```
 cd tekton && ./bootstrap.sh
 ```

This should complete successfully.

Wait for the `demo-interceptor` and `el-cicd-event-listener` pods to be up and running.

## Setup the Webooks

  For your GitHub repositories, go to the settings page, and add the route paths
that have been setup.
