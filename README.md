# Upload
Lazypic S3 스토리지에 파일(블랜더파일, 이미지파일)을 업로드하는 코드입니다.
애니메이션 작업 파일을 Upload, Publish 를 위해 사용합니다.

#### 사용법
- 홈디렉토리 render.blend 파일을 pilot프로젝트, 에피소드1, 씬1, 컷2 파일로 S3 클라우드에 올리는 예제.
```bash
$ upload -project pilot -ep 1 -s 1 -c 2 -file ~/render.blend
```

#### 파일 버전관리
- S3 기능중 `속성 > 버전관리`를 활성화하면 같은 key로 매번 파일업로드를 하더라도 자동으로 버전관리가 되어 따로 코드에서 버전관리를 할 필요가 없습니다.

#### License
- AWS API는 Apache License 2.0 라이센스 정책을 따르기 때문에 이 코드도 같은 라이센스를 사용합니다.
