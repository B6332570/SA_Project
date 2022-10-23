import { ProvincesInterface } from "./IProvince";
import { RolesInterface } from "./IRole";
import { MemberClassesInterface } from "./IMemberClass";
import { EmployeesInterface } from "./IEmployee";
export interface UsersInterface {
    ID?:        number,
    Pin?:       string,
	FirstName?: string,
	LastName?:  string,
	Civ?:       string,
	Phone?:     string,
	Email?:     string,
	Password?:  string, 
	Address?:   string,
	ProvinceID?:    number,
	Province?:    ProvincesInterface,  
	RoleID?:        number,    
	Role?:        RolesInterface,       
	MemberClassID?: number, 
	MemberClass?: MemberClassesInterface, 
	EmployeeID?:    number,    
	Employee?:    EmployeesInterface,    
}