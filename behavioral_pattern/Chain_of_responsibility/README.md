# Chain of responsibility
# Khi nào dùng:
1. Khi chúng ta có 1 đối tượng mà chúng ta phải xử lý 1 chuỗi các hành vi xữ lý tuần tự (Ta có 1 đối tượng khi ta xử lý hành vi A xong chúng ta xử lý  hành vi B xong rồi sang c,...)

# Các thành phần:
1. Clients: Cái đối tượng mà chúng ta cần xử lý theo chuỗi các hành vi  
2. Handler interface: mà chúng ta sẽ định nghĩa các hành vi chung cho chuỗi các công việc mà ta thực hiện.
    + Execute : Mỗi hành vi chúng ta sẽ thực hiện hành động j đó
    + SetNext: Chúng ta sẽ cài đăt các hanh vi kế tiếp trong chuỗi hành vi
3. concrete handler 1
4. concrete handler 2
5. concrete handler 3
6. concrete handler 4

# Ví dụ: 
Quá trình khám bệnh của bệnh nhân.
1. patient
2. Department
3. Reception
4. Doctor
5. Medicia  
6. Cashier