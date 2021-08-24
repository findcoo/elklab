- [읽기전](#읽기전)
- [분석 개요](#분석-개요)
- [분석 과정](#분석-과정)
  - [분석이 작동하는 시점](#분석이-작동하는-시점)
  - [색인 분석시에서 어떤 분석기가 사용될까](#색인-분석시에서-어떤-분석기가-사용될까)
  - [분석기 설정의 유연성](#분석기-설정의-유연성)
    - [질의와 색인 분석기를 다르게 설정하는 이유](#질의와-색인-분석기를-다르게-설정하는-이유)
- [어근 추출(Stemming)](#어근-추출stemming)
- [Token graph](#token-graph)
- [Normalizer](#normalizer)

### 읽기전

본 문서에 예제들은 js를 기반으로 작성한 의사코드 입니다.

### 분석 개요

데이터를 색인하는 과정에서 검색어를 추출하기 위한 과정을 거치는데, 이를 **분석**이라한다.
이 과정에서 사용되는 프로그램을 분석기(analyzer)라 하며 하나의 `tokenizer`와 0개 이상의 `token filter`로 구성된다.

공식문서에서는 추가적으로 `Character filter`에 대해 설명한다.
문장을 문자단위로 변경할 때 사용된다.

```js
const hinduNumbers = '٠‎١٢٣٤٥٦٧٨‎٩'  

const arabicNumbers = arabicNumberCharacterFilter(hinduNumbers)
// arabicNumbers = 0123456789
```

### 분석 과정

tokenizer -> token filter

문자가 위 순서대로 tokenizer와 token filter를 거쳐 검색어로 재정렬된다.

```js
let text = 'Around the World in Eighty Days'

const tokens = whitespaceTokenizer(text)
// tokens = ['Around', 'the', 'World', 'in', 'Eighty', 'Days']

const analyzedText = stopFilter(lowercaseFilter(tokens))
// analyzedText = ['around', 'world', 'eighty', 'days']
```

#### 분석이 작동하는 시점

분석은 두가지 경우에 발생한다.

* 문서(document) 색인시 모든 `text` 필드의 값은 분석된다.
* `text` 필드에 대해 전문 검색시에 질의된 문장은 분석된다.

#### 색인 분석시에서 어떤 분석기가 사용될까

색인용 분석기를 선택함에 있어 두가지 순서에 따라 결정된다.

* `mapping`에서 필드에 설정한 분석기를 선택한다.
* `analysis.analyzer.default` 색인 설정에 분석기를 선택한다.

#### 분석기 설정의 유연성

앞서 색인과 검색시점에 분석이 사용됨을 알았다면
이번에는 색인과 검색에 사용되는 분석기가 서로 달라질수도 있음을 알아보자.
이는 분석기 설정이 유연하게 설계되어 가능하다.

질의시에 분석기를 설정한다.

```js
GET my-index-000001/_search
{
  "query": {
    "match": {
      "message": {
        "query": "Quick foxes",
        "analyzer": "stop"
      }
    }
  }
}
```

필드에 대해서 분석기 설정도 가능하다.
```js
PUT my-index-000001
{
  "mappings": {
    "properties": {
      "title": {
        "type": "text",
        "analyzer": "whitespace",
        "search_analyzer": "simple"
      }
    }
  }
}
```

인덱스 기본 분석기 설정도 가능하다.

```js
PUT my-index-000001
{
  "settings": {
    "analysis": {
      "analyzer": {
        "default": {
          "type": "simple"
        },
        "default_search": {
          "type": "whitespace"
        }
      }
    }
  }
}
```

위 세가지 분석기 설정을 통해 질의시와 색인시의 분석기를 다르게 설정할수 있음을 알수있다.

##### 질의와 색인 분석기를 다르게 설정하는 이유

색인을 통해 **apple**이라는 단어를 색인하여
`['a', 'ap', 'app', 'appl', 'apple']` 이라는 토큰들로 색인되었다고 가정해보자.

사용자는 **appli**라는 문장으로 질의를 진행한다.
동일한 분석기를 사용하면 `['a', 'ap', 'app', 'appl', 'appli']` 라는 토큰으로 분석될것이다.
4개의 토큰이 매칭됨으로 **apple**은 질의에 의해 검색될것이다.

여기서 문제는 사용자는 a나 ap, app을 포함하는 문장을 찾고자하기보단 **appli**를 포함한 문장을 찾고 싶어
질의를 했을것인데 의도하지 않는 결과가 더 많이 검색된다.

사용자의 질의의도에 맞게 검색결과를 노출하기위해 질의와 색인의 분석기를 다르게 설정하여 사용하는 요구사항이 발생한다.

**요약**

```
사용자: appli 로 시작하는 단어는 뭐가 있을까?
검색결과: apple
사용자: 왜 이게 나와?
```

### 어근 추출(Stemming)

stem이라는 단어는 줄기를 제거한다는 뜻이있다.
**walking**과 **walked**라는 단어들이 있으면 어근이 되는 **walk** 토큰으로 추출한다. 파생어에서 어근을 추출하는 것을 **Stemming**이라 칭한다.

흔히 말하는 형태소 분석기는 이 어근 추출을 위한 `Stemmer token filter`를 뜻한다. 형태소는 뜻을 가진 가장 작은 말의 단위를 뜻하며 형태소도 여러 종류가 있다. 이중에서 어근은 실질형태소에 포함되며
파생어(실질 형태소에 접사가 결합된 단어)에서 어근을 추출하는 것을 Stemming이라 정의할수있다.

두 가지 종류의 Stemmer가 있다.

* 알고리즘 Stemmer, 어근을 추출할때 규칙을 기준으로 추출
* 사전형 Stemmer, 사전에 등록된 어근을 기준으로 추출

### Token graph

tokenizer는 문장을 토큰으로 변경한다. 토큰을 생성하면서 부수적으로
생성되는 것이 있다.

* 토큰열에 토큰들의 위치
* 토큰의 길이

즉, 토큰은 루프구조가 없는 방향성 비순환 그래프로 구성된다(DAG).
토큰의 위치 정보를 통해 동의어 및 약어를 토큰으로 구성할 수 있다.

```js
/**
 *  
 *  0 --- quick --> 1 --- brown --> 2 --- fox --> 3
 *  |               ^
 *  |               |
 *  fast ------------
 */

const tokens = [ ['quick', 'fast'], 'brown', 'fox' ]
```

quick과 fast 그래프상에 같은 위치임으로 동의어다.

```js
/**
 *   0 --- domain --> 1 --- name --> 2 --- system --> 3
 *   |                                                ^
 *   |                                                |
 *   dns ----------------------------------------------
 */ 

const tokens = [ ['domain', 'dns'], 'name', ['system', 'dns'] ]
```
dns는 domain name system의 약어이다.


### Normalizer

분석기와 비슷하지만 차이점으로 토큰을 생성하지 않는다. 오직 하나의 문장으로 변경한다. 문자단위로만 작동하는 token filter와 char filter만 사용할 수 있다.





