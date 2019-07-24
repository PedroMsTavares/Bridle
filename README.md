# Bridle

 _Bridle_: A system used to share the strain of an anchor evenly across two points. It reduces wear on the yacht, acts as a shock absorber, and silences the rattle of the anchor chain. Also helps prevent your catamaran from floating away

## Requirement:

Create a solution to copy images stored externaly, to internal repositories, so in case of external failure we can switch to internal ones. 

## How it's works ?

_Bridle_ is tiny program writen in Go, that receives a paylod from a Admission Webwook every time a new pod is created. Parses the payload and validates if its a internal image (ecr) and if not copys the image to a internal ECR.  At the moment it only support public images to ECR.


## Deployment: 

Run `make deploy` this assumes you have helm. 
    
