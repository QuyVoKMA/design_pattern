# Tạm dich là giao diện bên ngoaì
#Khi nào dùng:
1. Khi chúng ta làm việc số lượng lơn các đối tượng với nhiều
    cách khỏi tạo khác nhau, và cách kết hợp hành vi của chúng phức tạp
    Chúng ta ko muốn người sd và bên thứ 3 phải trực tiếp làm việc vơi giao diện đó
    ==> cung cấp facade: là giao diện dễ sử dụng hơn cho nguoif sd.
    Người sd sẽ ko cần phải quan tâm đến cái phức tạp bên trong mà chúng ta sẽ xử lý lớp bên trong này.

# Thành phần:
-Subsystem 1 -->account

-Subsystem 2 -->securityCode

-Subsystem 3 -->wallet

-Subsystem 4 -->notification

==>facade -->wallet

==> Mỗi subsystem này là một đối tượng và cách khỏi tạo cũng khách nhau
    các hành vi khác nhau tuy nhiên nó tạo ra 1 hệ thống để phục vụ cho mục đích nào đó

-Người sđ sẽ ko biến các thành phần bên trong là gì, người ta chỉ biến nó là FACADE
    
FACADE là một đối tượng chứa tất cả các đối tượng con, been trong có các hành vi bên trong hvi này nó sẽ chứa tất cả các
    hành vi của các đối tượng con