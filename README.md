# TODO List in Golang

## Methods
* GetTodoList : TODO 목록 조회
* AddTodo : TODO  등록
* UpdateTodo : TODO 상태 수정
* DeleteTodo : TODO  삭제

## Databases
* sqlite3 사용.
* table
    * tb_todo
        * id : seq coloms, pk, auto increment
        * title : todo titles, text, not null
        * doneYN : 진행상태, boolean, not null
        * createDt : 생성일자, varchar(14), not null
        * updateDt : 수정일자, varchar(14), not null
    * Create SQL

    ```
        
        CREATE TABLE tb_todo (
            id integer primary key autoincrement, 
            title text not null,
            doneYN boolean not null,
            createDt varchar(14) not null,
            updateDt varchar(14) not null
        )
    ```
        