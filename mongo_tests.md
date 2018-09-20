# Queries

## Find movies by "Stanley Kubrick"

{"director" : "Stanley Kubrick"}

## That where released in 1987

{"director" : "Stanley Kubrick", "year": 1987}

## Exclude 1960

{"director" : "Stanley Kubrick", "year" : {"$ne": 1960}}

## This is equivalent

{"director" : {"$eq" : "Stanley Kubrick"}, "year" : {"$ne": 1960}}

## Should return nothing

{"director" : "Stanley Kubrick", "year" : {"$ne": 1960}, "title" : "Spartacus"}

## Are there other Kubricks?

{"director" : {$regex : ".*Kubrick.*"} }

## Reliably rated classics

{"viewerVotes" : {"$gt" : 2000}, "viewerRating" : {"$gt" : 9.0} }

## Best avoided

{"viewerVotes" : {"$gt" : 2000}, "viewerRating" : {"$lt" : 5.0} }

## All Clint Eastwood movies

{"cast": "Clint Eastwood"}


