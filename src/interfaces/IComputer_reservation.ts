import { ComputersInterface } from "./IComputer"
import { Computer_ossInterface } from "./IComputer_os"
import { Time_comsInterface } from "./ITime_com"
import { UsersInterface } from "./IUser"

export interface Computer_reservationInterface{
	ID?: Number;
	Date?: Date | null;

	Computer_id?: number;
	Computer?: ComputersInterface;

	Computer_os_id?: number;
	Computer_os?: Computer_ossInterface;

	Time_com_id?: number;
	Time_com?: Time_comsInterface;

	UserID?: number;
	User?: UsersInterface;
}