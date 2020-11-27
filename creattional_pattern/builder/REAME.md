# Builder 

# Khi nào sữ dụng:

1. Khi chúng ta muốn khỏi tạo một đối tượng của struct nhưng struct đó có rất nhiều thuộc tính, khi chúng ta khỏi tạo một cách bình thường chúng ta phải nắm được kiểu dữ liệu của thuộc tính đó, thứ tự của thuộc tính đo, và cái cách chung ta truyền vào cho chính xác
2. Vd ta có 1 struct account tài khoảng ngân hàng, trong 1 tài khoảng có rất nhiều thuộc tính như tên, id, địa chỉ, giới tính, số tiền nợ, số tiền nợ, số tiền lãi,...Như vậy khi chúng ta sd phải gọi hàm khỏi tạo, chúng ta phải sắp xếp đúng thứ tự truyền vào. Trong trường hợp xấu nhất, chúng ta có những thuộc tính có cùng kiểu dữ liệu khác nhau, chẳn hạng như số tuổi và số tiền đều là số nguyên, khi truyền và chúng ta lỡ truyền nhầm vào thì gây ra lỗi rất lớn
3. Khi chúng ta có 1 struct, chúng ta muốn tạo các đối tượng cụ thể hơn cho struct 
4. Khi chúng ta có 1 struct chúng ta khỏi tạo một đối tượng với đầy đủ các thuộc tính, hoặc không khỏi tạo các đối tượng. Ví dụ chúng ta có 1 struct có 4 thuộc tính thì chúng ta phải khỏi tạo đầy đủ 4 thuộc tính

# Các thành phần.

1. Product (Cái đối tượng chúng ta muốn tạo ra là đối tượng gì)
2. Builder interface   Chứa các phương thức gán các giá trị để gán các giá trị cho product này
3. Concrete builder 1
4. Concrete builder 2
5. Concrete builder 3 (Mỗi 1 cái này sẽ tạo ra một đối tượng cụ thể cho product này)
6.Director (Giúp chúng ta liên kết được giữa cái builder interface và nhiều cái concrete  builder của nó)

# Ví dụ
1. House
2. 
3. Nomal House
4. Igloo House
5
6