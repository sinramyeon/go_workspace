# 개발서적 정리용 사이트

---

## 주말에 golang과 siris(iris) 프레임워크 공부 겸 연습 겸 해서 제작
## 20170826 ~

## 사용할 프레임워크들

종류 | 내용
---- | ----
언어 | Go
웹 프레임워크 | siris
DB | MariaDB(또는 후에 Redis 이관)
차트 라이브러리 | http://gionkunz.github.io/chartist-js/
---
> 아이디어

1. 언어별, 분야별 진로를 만들어서 그 진로를 설정하고 추천 책 서비스나 얼만큼 어느 진로로 가는 분야를 공부했는지 보여주는 기능?
2. 그 진로에 따른 캐릭터를 만들어서 지금 당신의 아바타로 활용?
3. 디벨로퍼 로드맵 기능을 저런식으로 구현해도 재미있을듯?
3. 프로그래밍 질답 자유 게시판?
4. 지금 사람들이 읽고 있는 책, 최신 책, 인기 책 피드?
5. 바로 구매하기(알라딘으로 이동)
6. 네이버, 페이스북, 트위터, 깃허브로 가입 ( 알라딘, 교보문고, 예스24로 가입도 가능한가? )
7. 책 공부 진행률 넣기?(단원 단위? 그냥 사용자 입력 단위?)
8. 사용자로 로그인 시 메인화면에 책 선반이 뜨고 읽은 책들이 쭉 늘어놓아지는 것?

---
### 진행

- [ ] 메인화면
http://www.free-css.com/free-css-templates/page214/wesoon

- [ ] CSS, HTML UI
- [ ] 헤더, 메뉴, 푸터

---

- [ ] 사용자
- [ ] 회원가입, 회원수정, 회원탈퇴
- [ ] 로그인, 로그아웃
- [ ] oAuth를 활용한 네이버, 페북, 깃허브 등을 통한 회원가입

---

- [ ] 책 메뉴
- [ ] 읽은 책 CRUD

기능 | 설명
---- | ----
책 제목 | 책 제목을 입력 시 알라딘에서 검색해서 갖고온다
책 설명 | 입력 시 입력내용, 미입력 시 알라딘에서 겁색해서 갖고온다
언어 종류 | 파이썬, 자바... 등 의 분야를 찾기
(구상중...)

- [ ] 책 통계

기능 | 설명
---- | ----
언어 종류 | 파이썬, 자바.. sql, 인프라 등 분야별 읽은책 개수 통계
달별, 년별 | 달별, 년별 읽은책 개수 통계
쓴 돈 | 알라딘에서 책 원가를 가져와다 쓴 돈을 통계를 내준다
(구상중...)

---

- [ ] AWS에 올리기
(AWS? Azure? 아님그냥 집 서버?)

- [ ] 도메인 구입과 연결

---

- [ ]  DB(MariaDB)

> 차후 Redis 공부하려면 Redis 이관도 생각해 볼것

1. 사용자(사용자 계정정보)

```
CREATE TABLE `USER` (
	`USER_NO` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT ='자동증감용 인덱스', 
	`USER_TYPE` ENUM('IT','NO') NULL DEFAULT 'IT' COMMENT 'IT:IT종사직, NO:일반인!',
	`ROLE` ENUM('MEMBER','MANAGER') NOT NULL DEFAULT 'MEMBER' COMMENT ='유저 권한',
	`ID` VARCHAR(10) NOT NULL DEFAULT '',
	`PASS` VARCHAR(500) NOT NULL DEFAULT '',
	`NAME` VARCHAR(20) NOT NULL DEFAULT '',
	`HP` VARCHAR(13) NOT NULL DEFAULT '',
	`CREATED` DATETIME NOT NULL,
	PRIMARY KEY (`USER_NO`)
)
COMMENT='유저 계정 정보'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;
```

2. 책

```
CREATE TABLE `BOOK` (
	`BOOK_NO` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT ='자동증감용 인덱스', 
	`ID` VARCHAR(10) NOT NULL DEFAULT '',

	(구상중)

	`CREATED` DATETIME NOT NULL,
	PRIMARY KEY (`BOOK_NO`)
)
COMMENT='책 정보'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;
```