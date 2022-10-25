import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { Computer_reservationInterface } from "../interfaces/IComputer_reservation";
import { GetComputer_reservations } from "../services/HttpClientService";
import moment from "moment";

function Computer_reservations() {
  const [computer_reservations, setComputer_reservations] = useState<Computer_reservationInterface[]>([]);

  useEffect(() => {
    getComputer_reservations();
  }, []);

  const getComputer_reservations = async () => {
    let res = await GetComputer_reservations();
    if (res) {
      setComputer_reservations(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 50 },
    {
      field: "User",
      headerName: "ผู้ใช้",
      width: 150,
      valueFormatter: (params) => params.value.FirstName,
    },
    {
      field: "Date",
      headerName: "วันที่",
      width: 150,
      valueFormatter:(params) => moment(params.value.Date).format("DD/MM/YYYY")
    },

    {
      field: "Computer",
      headerName: "เครื่องคอมพิวเตอร์",
      width: 160,
      valueFormatter: (params) => params.value.Comp_name,
    },


    { field: "Time_com",
     headerName: "เวลาจอง",
      width: 150,
      valueFormatter: (params) => params.value.Time_com_period, },

    // { field: "Computer.Computer_os",
    // headerName: "ระบบปฏิบัติการ",
    // width: 150,
    // valueFormatter: (params) => params.value.Name,   },
  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการจองเครื่องคอมพิวเตอร์
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/computer_reservations/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={computer_reservations}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={6}
            rowsPerPageOptions={[6]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Computer_reservations;
