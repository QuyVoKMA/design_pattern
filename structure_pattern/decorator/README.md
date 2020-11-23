#Khi nào dùng:
Khi có nhu cầu thêm một tính năng mới vào 1 đối tượng
có sẵn mà ko làm ảnh hưởng tới đối tượng này

# Cách dùng:
## 1. component interface --> IPizza (chứa hành vi chung của 1 đối tượng đã có sẵn)
### concrete componet 1 -->ChickentPizza
### concrete componet 2 -->-->TomatotPizza
==> Muốn mở rộng trong golang (ko có thừa kế) phải cài đặt thêm
###concrete decoretor 1 -->papperdecorator (chứa tính năng thêm của cc1)
        +component propeties
###concrete decoretor 2 -->cheeseDecorator
        +component propeties
        
        
        
# ==> Nó rất linh động khi thêm tính năm mới.