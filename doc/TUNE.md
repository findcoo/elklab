## TRICK

엔지니어링 관련 이슈

### 이슈 ###
* JVM 최적화
  * 커널 설정
    * max locked memory
    
      `ulimit -l unlimited`
        
      락 메모리 사이즈를 조정하여 JVM swap out 되는 것을 막는다.
    * vm.max\_map\_count
    
      `sysctl -w vm.max_map_count=262144`
      
      엘라스틱 서치가 사용하는 파일 시스템의 성능향상을 위한 설정
