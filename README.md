# Upload
Lazypic S3 스토리지에 파일(블랜더파일, 이미지파일)을 업로드하는 코드입니다.
애니메이션 작업 파일을 Upload, Publish 를 위해 사용합니다.

#### 다운로드 & 사용법
- [Download](https://github.com/lazypic/upload/releases)
- Lazypic은 기본적으로 `ap-northeast-2`(서울리전), `lazypic` S3 버킷을 사용합니다.
- 홈디렉토리 render.blend 파일을 pilot프로젝트, 에피소드1, 씬1, 컷2 파일로 S3 클라우드에 올리는 예제입니다.
```bash
$ upload -project pilot -ep 1 -s 1 -c 2 -file ~/render.blend
```

- 다른회사에 다른 지역이라면 아래처럼 사용자 인수를 지원합니다.
- `us-west-2`(미서부오레건), `companyname`의 bucket에 데이터를 넣는 예제는 아래와 같습니다.
```bash
$ upload -region us-west-2 -bucket companyname -project pilot -ep 1 -s 1 -c 2 -file ~/render.blend 
```

- 이 프로그램이 작동되기 위해서는 ~/.aws/credentials 파일에 아래코드가 존재해야합니다.
```
[default]
aws_access_key_id = 
aws_secret_access_key = 
```

#### 사용자컴파일
- Go가 설치되어있다면, 터미널을 통해서 아래처럼 설치할 수 있습니다.
```bash
$ go get -u github.com/lazypic/upload
```

#### 파일 버전관리
- AWS S3 기능중 `속성 > 버전관리`를 활성화하면 같은 key로 매번 파일업로드를 하더라도 스토리지 레벨에서 자동 버전관리가 됩니다.
- 위 특징으로 인해서 따로 코드에서 버전관리를 할 필요가 없습니다.

#### License
- AWS API는 Apache License 2.0 라이센스 정책을 따르기 때문에 이 코드도 같은 라이센스를 사용합니다.
