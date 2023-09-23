#TOC
- [1. Formating](#1-formating)
- [2. Table](#2-table)
- [3. Code style](#3-code-style)
- [4. Hình ảnh](#4-hình-ảnh)
- [5. Vẽ biểu đồ](#5-vẽ-biểu-đồ)
  - [5.1. Mermaid](#51-mermaid)
  - [5.2. PlantUML](#52-plantuml)
- [6. Công thức toán](#6-công-thức-toán)


## 1. Formating

**đây là bôi đậm**
*In nghiêng*
~~Gạch ngang~~
> Đây là một `quotation`

## 2. Table
| #   | Title 1   | Title 2   |
| --- | --------- | --------- |
| 1   | Content 1 | Content 2 |
| 2   | Content 1 | Content 2 |

## 3. Code style

```go
func main() {
    fmt.Println(123)
}
```

```js
console.log(123)
```

## 4. Hình ảnh
![](1.png)

## 5. Vẽ biểu đồ

### 5.1. Mermaid

```mermaid
graph TB
    sq[Square shape] --> ci((Circle shape))

    subgraph A
        od>Odd shape]-- Two line<br/>edge comment --> ro
        di{Diamond with <br/> line break} -.-> ro(Rounded<br>square<br>shape)
        di==>ro2(Rounded square shape)
    end

    %% Notice that no text in shape are added here instead that is appended further down
    e --> od3>Really long text with linebreak<br>in an Odd shape]

    %% Comments after double percent signs
    e((Inner / circle<br>and some odd <br>special characters)) --> f(,.?!+-*ز)

    cyr[Cyrillic]-->cyr2((Circle shape Начало));

     classDef green fill:#9f6,stroke:#333,stroke-width:2px;
     classDef orange fill:#f96,stroke:#333,stroke-width:4px;
     class sq,e green
     class di orange
```

### 5.2. PlantUML
```plantuml
@startuml
participant Participant as Foo
actor       Actor       as Foo1
boundary    Boundary    as Foo2
control     Control     as Foo3
entity      Entity      as Foo4
database    Database    as Foo5
collections Collections as Foo6
queue       Queue       as Foo7
Foo -> Foo1 : To actor 
Foo -> Foo2 : To boundary
Foo -> Foo3 : To control
Foo -> Foo4 : To entity
Foo -> Foo5 : To database
Foo -> Foo6 : To collections
Foo -> Foo7: To queue
@enduml
```

## 6. Công thức toán
$y = x^2$


Ex1: \w+.(jpg|png|jpeg)$
Ex2: ^[0-9.,]+$
Ex3: ^[\w.]+@[\w-]+(\.\w+)+$
Ex4: interface (?<interface>.+)\sstatus (?<status>\w+)\sip address (?<ip00>[0-9. ]+)\s(ip address secondary (?<ip01>[0-9. ]+))?


Step 1: ^((?!voucher).)*$
Step 2:   INFO.+User
Step 3:  claimed voucher