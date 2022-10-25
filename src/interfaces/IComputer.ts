import { Computer_ossInterface } from "./IComputer_os";

export interface ComputersInterface{
	ID: number,
    Comp_name: string,
	Comp_room: string,
	Computer_os_id: number,
	Computer_os?: Computer_ossInterface;
}