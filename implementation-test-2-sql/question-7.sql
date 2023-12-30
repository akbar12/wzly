SET @row_number = 0; 
select (@row_number:=@row_number + 1) AS rank, e.name, sum(s.sales) as sales
from sales_data s
left join employees e 
on e.employee_id = s.employee_id
group by s.employee_id
order by sales desc ;