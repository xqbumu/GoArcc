package ArchitectureSuggesterService

import (
	"alfred/logger"
	"alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/ArchitecureConfiguration"
	awsDefaultPb "alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/ArchitecureConfiguration/pb"
	"alfred/modules/ArchitectureSuggesterService/pb"
	"go.uber.org/zap"
)

type ArchitectureSuggesterService struct {
	awsDefaultData *awsDefaultPb.AwsDefaultData
}

//todo : AlWays add migration code for best practices
func NewArchitectureSuggesterService() pb.ArchitectureSuggesterServiceServer {
	data, err := ArchitecureConfiguration.ParserData("awsConfig.yml")
	if err != nil {
		logger.Log.Fatal("not able to register", zap.Error(err))
	}
	return &ArchitectureSuggesterService{
		awsDefaultData: data,
	}
}
