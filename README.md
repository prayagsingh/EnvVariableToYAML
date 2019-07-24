# EnvVariableToYAML
### Reading environment variables and putting it into YAML file
##### This script is designed to add custom organisation names in the docker-compose-e2e-template.yaml file(HLF). 
##### Input :
 ###### docker-compose-e2e-template.tml.yaml file 
 ###### set env variables like ORDERER_PROFILE, ORD_CHANNEL_ID, CHANNEL_PROFILE, CHANNEL_NAME, ORG1_NAME, ORG2_NAME, DOMAIN, PROJECT_NAME
 
##### Output : docker-compose-e2e-temp.yaml file
##### Bugs:
  ###### conversion from YAML to JSON add "null" in the final YAML file like "networks: null"
  ###### some content skipped in YAML to JSON conversion (JSON file started from networks:)
  
##### Working 95% 

