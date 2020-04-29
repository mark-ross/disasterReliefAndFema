# FEMA and Foursquare Disaster Relief

## Purpose
The goal of this project is to link with the Disaster Declaration API listed
on the FEMA website.  The `fema` folder contains an API definition and a
start of a collection of filters to better process the results on the
server side.

Once the FEMA reports are processed, the goal will be to locate Foursquare
churches that are close to the disaster and contact/email the respective
points of contact to alert them to the potential help they can receive.

## Why?
Foursquare Disaster Relief offers services and only require a quick contact
to begin the process of receiving resources to help the local area to begin
the recovery process of the community in pain.

However, there are instances where people don't know that they can call FDR.
Additionally, there may be people who want to help and can be trained by FDR.

## Status

### 2020-03-08

This project is in its early stages.  The foundational API of FEMA has been
'sketched' and some of the filter options available.  Specifying a date filter
fails to produce results, so more time is needed to debug this.

# FEMA USAGE REQUIREMENTS
This product uses the Federal Emergency Management Agency’s API, but is not endorsed by FEMA.

This product accessing the information from the API on a per request.  Therefore the date
associated with the information displayed will be the present date.  If any information is
archived, please be sure to display that information.

For more information, please see the 
[FEMA API Terms and Conditions]("https://www.fema.gov/openfema-api-terms-conditions").