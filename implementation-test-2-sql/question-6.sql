select e.name, e.salary, e2.avg_salary
from employees e
join (
	select department, AVG(salary) as avg_salary
	from employees
	group by department
	order by avg_salary desc limit 1
) as e2
on e.department = e2.department;