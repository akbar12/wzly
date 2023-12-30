select e.name, sum(s.sales) as sales 
from sales_data s 
left join employees e 
on e.employee_id = s.employee_id
group by s.employee_id
order by sales desc limit 5