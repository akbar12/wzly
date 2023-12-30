CREATE DEFINER=`root`@`localhost` PROCEDURE `search_employee`(
	IN department_filter varchar(45)
)
BEGIN
	select name, salary from employees where department = department_filter;
END;

CALL search_employee('sales');