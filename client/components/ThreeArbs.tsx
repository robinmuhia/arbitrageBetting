import { Box, useTheme } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import Header from "./Header";
import { useEffect, useState } from "react";
import ThreeArbModal from "./ThreeArbModal";

const ThreeArbs = () => {
  const [arbData, setArbData] = useState([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const theme = useTheme();
  const columns = [
    {
      field: "Title",
      headerName: "Game",
      sortable: false,
      filterable: false,
      flex: 1.2,
    },
    {
      field: "Home",
      headerName: "Home Team",
      sortable: false,
      filterable: false,
      flex: 0.6,
    },
    {
      field: "Draw",
      headerName: "Draw Team",
      sortable: false,
      filterable: false,
      flex: 0.6,
    },
    {
      field: "Away",
      headerName: "Away Team",
      sortable: false,
      filterable: false,
      flex: 0.6,
    },
    {
      field: "HomeOdds",
      headerName: "Home odds",
      sortable: false,
      filterable: false,
      flex: 0.5,
      renderCell: (params: any) => `${Number(params.value).toFixed(2)}`,
    },
    {
      field: "DrawOdds",
      headerName: "Draw odds",
      sortable: false,
      filterable: false,
      flex: 0.5,
      renderCell: (params: any) => `${Number(params.value).toFixed(2)}`,
    },
    {
      field: "AwayOdds",
      headerName: "Away odds",
      sortable: false,
      filterable: false,
      flex: 0.5,
      renderCell: (params: any) => `${Number(params.value).toFixed(2)}`,
    },
    {
      field: "GameType",
      headerName: "Sport",
      sortable: false,
      filterable: false,
      flex: 0.5,
      renderCell: (params: any) =>
        `${params.value.split("_")[0].toUpperCase()}`,
    },
    {
      field: "League",
      headerName: "League",
      sortable: false,
      filterable: false,
      flex: 0.5,
    },
    {
      field: "Profit",
      headerName: "Profit",
      sortable: false,
      filterable: false,
      flex: 0.5,
      renderCell: (params: any) => `${Number(params.value).toFixed(2)}%`,
    },
    {
      field: "GameTime",
      headerName: "Game Time",
      flex: 0.8,
      filterable: false,
      renderCell: (params: any) =>
        `${params.value.slice(11, 16)} ${params.value.slice(
          8,
          10
        )}/${params.value.slice(5, 7)}/${params.value.slice(0, 4)}`,
    },
    {
      field: "",
      headerName: "Calculate arbs",
      sortable: false,
      filterable: false,
      flex: 0.8,
      renderCell: (params: any) => <ThreeArbModal props={params.row} />,
    },
    {
      field: "BookmarkerRegion",
      headerName: "Bookmarkers Location",
      sortable: false,
      filterable: false,
      flex: 0.5,
      renderCell: (params: any) => `${params.value.toUpperCase()}`,
    },
  ];

  useEffect(() => {
    const fetchArbs = async () => {
      try {
        const url = process.env.NEXT_PUBLIC_BACKEND_URL;
        const res = await fetch(`${url}/api/v1/threearbsbets`, {
          method: "GET",
          credentials: "include",
        });
        if (res.ok) {
          const data = await res.json();
          const { threeArbs } = data;
          setArbData(threeArbs);
          setIsLoading(false);
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };
    fetchArbs();
  }, []);

  return (
    <Box m="1.5rem 2.5rem">
      <Header
        title="Three Arbitrage Bets"
        subtitle="List of all Current Available for three arbitrage bets"
      />
      <Box
        height="80vh"
        marginTop="10px"
        sx={{
          "& .MuiDataGrid-root": {
            border: "none",
          },
          "& .MuiDataGrid-cell": {
            borderBottom: "none",
          },
          "& .MuiDataGrid-columnHeaders": {
            //@ts-ignore
            backgroundColor: theme.palette.background.alt,
            //@ts-ignore
            color: theme.palette.secondary.main,
            borderBottom: "none",
          },
          "& .MuiDataGrid-virtualScroller": {
            backgroundColor: theme.palette.primary.main,
          },
          "& .MuiDataGrid-footerContainer": {
            //@ts-ignore
            backgroundColor: theme.palette.neutral.main,
            color: theme.palette.secondary.main,
            borderTop: "none",
          },
          "& .MuiDataGrid-toolbarContainer .MuiButton-text": {
            //@ts-ignore
            color: `${theme.palette.primary.main} !important`,
          },
          "& .MuiDataGrid-panelFooter .css-4rdffl-MuiDataGrid-panelFooter": {
            backgroundColor: theme.palette.background.default,
            //@ts-ignore
            color: theme.palette.secondary[100],
          },
        }}
      >
        <DataGrid
          loading={isLoading || !arbData}
          //@ts-ignore
          getRowId={(row) => row.ID}
          rows={arbData || []}
          columns={columns}
        />
      </Box>
    </Box>
  );
};

export default ThreeArbs;
