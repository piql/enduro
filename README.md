
<p align="left">
  <a href="https://github.com/penwern/enduro/releases/latest"><img src="https://img.shields.io/github/v/release/artefactual-labs/enduro.svg?color=orange"/></a>
  <img src="https://github.com/penwern/enduro/workflows/Test/badge.svg"/>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache%202.0-blue.svg"/></a>
  <a href="https://goreportcard.com/report/github.com/penwern/enduro"><img src="https://goreportcard.com/badge/github.com/penwern/enduro"/></a>
  <a href="https://codecov.io/gh/artefactual-labs/enduro"><img src="https://img.shields.io/codecov/c/github/artefactual-labs/enduro"/></a>
</p>

# Enduro

Enduro is a tool designed to automate the processing of transfers in multiple
Archivematica pipelines. It's part of a preservation solution that is being
used by the [National Health Archive (NHA)] and the [National Center for Truth
and Reconciliation (NCTR)].

It's a **proof of concept** at its very early stages. It aims to cover our
client's needs while exploring new and innovative ways to build durable and
fault-tolerant workflows suited for preservation.


## E-Ark AIP Generator Workflow

### Original Development Installation Instructions
https://enduroproject.netlify.app/docs/development/environment/

### Additional Installation Instructions 
To clone this branch from scratch use:
`git clone --recurse-submodules https://github.com/penwern/enduro.git`
Otherwise the submodules will need to be pulled:
`git submodule update --init --recursive`

You will need to edit the **enduro.toml** `[pipeline]` section with your Archivematica details


### Additional Dependencies
Python 3.9	-	for scripts
 - python3.9 pip - `curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py` and `python3.9 get-pip.py`
	 - xmltodict - `python3.9 -m pip install xmltodict`
	 - metsrw - `python3.9 -m pip install metsrw`
		 - Error: python setup.py egg_info did not run successfully.
		 - Solution: `python3.9 -m pip install --upgrade setuptools`

Java 8	- for Commons IP Validator

### Execution
Once Enduro has been set up you will be able execute the workflow.

SIPs to be processed through the workflow should be placed in the **sips/** directory.
Currently, SIPs are not deleted from the **sips/** directory after processing.
Resulting E-Ark AIPs will be found in the **eark_aips/** directory.
Validation reports for both can be found in their respective directories.

Example command to run the workflow:
`./hack/cadence.sh wf start --wid="eark-aip-generator" --wt="eark-aip-generator" --wrp="1" --tl="global" --et="3600"`

### API

We have extended the API service to include the service E-Ark. Currently, this service executes our eark-aip-generator workflow, creating E-Ark compliant AIPS, using SIPs in the sips/ directory and outputting into eark_aips/.

POST /eark
 - Runs the workflow.
 
GET /eark
 - Gets the status of the running/last run workflow.

## Further reading

Visit https://enduroproject.netlify.com for more details.

[National Health Archive (NHA)]: https://www.piql.com/norwegians-digital-health-data-to-be-preserved-for-future-generations/
[National Center for Truth and Reconciliation (NCTR)]: https://nctr.ca/about/about-the-nctr/our-mandate/
