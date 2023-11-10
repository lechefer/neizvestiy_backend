package config

type Config struct {
	Env       Env
	Port      int
	Databases Databases `json:"databases"`
	S3        S3        `json:"s3"`
}

type Databases struct {
	Postgres Database `json:"postgres"`
}

type Database struct {
	ConnectionString string `json:"connectionString"`
	MigrationPath    string `json:"migrationPath"`
}

type S3 struct {
	Endpoint    string  `json:"endpoint"`
	Buckets     Buckets `json:"buckets"`
	AccessKeyId string  `json:"accessKeyId"`
	SecretKey   string  `json:"secretKey"`
}

type Buckets struct {
	Main Bucket `json:"Main"`
}

type Bucket struct {
	Name         string `json:"name"`
	PublicDomain string `json:"publicDomain"`
}
