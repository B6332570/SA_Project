import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";

import { ComputersInterface } from "../interfaces/IComputer";
import { Computer_ossInterface } from "../interfaces/IComputer_os";
import { Time_comsInterface } from "../interfaces/ITime_com";
import { UsersInterface } from "../interfaces/IUser";
import { Computer_reservationInterface } from "../interfaces/IComputer_reservation";


import {
    GetComputers,
    GetComputer_oss,
    GetTime_coms,
    GetUsers,
    Computer_reservations,
} from "../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function Computer_reservationCreate() {
  const [computers, setComputers] = useState<ComputersInterface[]>([]);
  const [time_coms, setTime_coms] = useState<Time_comsInterface[]>([]);
  const [users, setUsers] = useState<UsersInterface[]>([]);
  const [computer_oss, setComputer_oss] = useState<Computer_ossInterface[]>([]);
  const [computer_reservation, setComputer_reservation] = useState<Partial<Computer_reservationInterface>>({
    Date: new Date(),
  });

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (event: SelectChangeEvent) => {
    console.log(event.target.name);
    console.log(event.target.value);

    
    const name = event.target.name as keyof typeof computer_reservation;
    setComputer_reservation({
      ...computer_reservation,
      [name]: event.target.value,
    });
  };
  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof computer_reservation;
    const { value } = event.target;
    setComputer_reservation({ ...computer_reservation, [id]: value });
  };

  const getComputers = async () => {
    let res = await GetComputers();
    if (res) {
      setComputers(res);
    }
  };

  const getComputer_oss = async () => {
    let res = await GetComputer_oss();
    if (res) {
      setComputer_oss(res);
    }
  };

const getTime_coms = async () => {
    let res = await GetTime_coms();
    if (res) {
      setTime_coms(res);
    }
  };

const getUsers = async () => {
    let res = await GetUsers();
    if (res) {
      setUsers(res);
    }
  };

  

  useEffect(() => {
    getComputers();
    getComputer_oss();
    getTime_coms();
    getUsers();
    
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      Computer_id: convertType(computer_reservation.Computer_id),
      Computer_os_id: convertType(computer_reservation.Computer_os_id),
      Time_com_id: convertType(computer_reservation.Time_com_id),
      UserID: convertType(localStorage.getItem("id") as string),
      Date: computer_reservation.Date,
    };

    let res = await Computer_reservations(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  return(
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Computer reservation
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 3 }}>

        <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>OS</p>
              <Select
                native
                value={computer_reservation.Computer_os_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Computer_os_id",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือก OS
                </option>
                {computer_oss.map((item: Computer_ossInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid> 
          
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Computer</p>
              <Select
                native
                value={computer_reservation.Computer_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Computer_id",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเครื่องคอมพิวเตอร์
                </option>
                {computers.map((item: ComputersInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Comp_name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>


          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={computer_reservation.Date}
                  onChange={(newValue) => {
                    setComputer_reservation({
                      ...computer_reservation,
                      Date: newValue,
                    });
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Time</p>
              <Select
                native
                value={computer_reservation.Time_com_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Time_com_id",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเวลาที่ต้องการจอง
                </option>
                {time_coms.map((item: Time_comsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Time_com_period}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/computer_reservations"
              variant="contained"
              color="inherit"
            >
              ย้อนกลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );

}

export default Computer_reservationCreate;