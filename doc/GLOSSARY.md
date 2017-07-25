# ElasticSearch Glossary
  
  엘라스틱 서치의 주요 용어 정리

### Terms ###
* analysis
  * fulltext를 단어로 변형하는 작업.
* cluster
  * 수개의 노드로 이루어진 군집. 
    * 노드는 엘라스틱 서치 데몬(서버)를 뜻함
  * 하나의 클러스터에는 하나의 마스터가 있다.
  * 마스터 노드가 죽으면 나머지 노드들 중 하나가 마스터가 된다.
* document
  * json 형태로 저장된다.
  * 일반 RDB에서 table의 row역활을 한다.
  * 각각의 doucment는 index에 저장된다.
  * type과 id를 갖는다.
  * 0개이상의 필드값을 갖는다.
  * 색인된 document는 `_source` 필드에 저장된다.
    * 검색하거나 조회시 기본적으로 사용되는 필드
* id
  * document의 id
  * `index\type\id` 값은 document 끼리 유일한 값을 나타낸다.
  * 입력하는 document값에 id 가 없다면 자동생성 된다. 
* field
  * document는 field들로 이루어진다. 
  * field의 값은 데이터형을 담는다(ex: date, string, integer)
  * RDB에서 table의 column과 유사하다.
* index
  * RDB에서 table과 유사하다. 
  * mapping을 통해 index의 field를 정의한다.
  * primary shard와 replica shard에 mapping을 정의하기위한 이름 역활을 한다.
* mapping
  * RDB에서 schema 정의와 유사하다.
  * index의 범위와 index의 type을 정의한다.
  * document를 통해 암시적으로 mapping이 생성될 수 있다.
* primary shard
  * document들은 하나의 primary shard에 저장된다.
  * document를 색인할 때, primary shard가 먼저 색인되고 replica shard들이 색인된다.
  * 기본 설정상 index하나는 5개의 primary shard를 갖는다.
  * index가 생성되면 추후 primary shard 수를 변경 할 수 없다.
* replica shard 
  * 0개 이상 존재할 수 있다.
  * primary shard가 죽을 경우 승격을 통해 primary shard가 된다.
  * Get, Search의 분산 처리를 위해 사용된다.
  * primary shard와 replica shard가 하나의 node에 존재하지 않는다.
* routing
  * shard의 생성 기준을 정한다.
  * 기본 설정상 document의 id를 기준으로 한다.
  * 부모 자식 관계의 document가 있을시 상관 관계에 의해 생성된다.
  * 색인시, mapping 시 정의 할 수 있으며 설정이 갱신될 수 있다.
* shard
  * 하나의 lucene instance다.
* source field
  * 색인된 document들은 `_source` 필드에 저장된다.
  * 검색을 통해 object를 찾기 위한 필드이다.
* term
  * 실제 색인된 데이터
  * term query를 통해 검색된다.
* text
  * 구조화 되기전 순수한 문장
  * term으로 분석될 수 있는 대상
* type
  * document의 type을 뜻한다.(ex: email, tweet, user)
  * 검색 api type을 통해 document를 구분한다.
  * index는 복수의 type을 갖는다.
  * type은document를 정의하는 복수의 field를 갖는다.
  * 하나의 index에서 다른 type의 같은 이름의 field는 똑같은 mapping을 갖는다.
