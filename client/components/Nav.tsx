"use client";

import { useState } from "react";
import { AppBar, Box, Toolbar, Typography, Button } from "@mui/material";

const Nav = () => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(true);

  const handleClick = async () => {
    const url = process.env.NEXT_PUBLIC_BACKEND_URL;
    const response = await fetch(`${url}/api/v1/logout`, {
      credentials: "include",
    });
    console.log(url);
    if (response.ok) {
      setIsLoggedIn(false);
    }
  };

  return (
    <>
      <Box
        sx={{
          display: "flex",
          flexGrow: 1,
          justifyContent: "center",
          alignItems: "center",
          boxSizing: "border-box",
        }}
      >
        <AppBar position="relative" color="primary">
          <Toolbar variant="dense">
            <Box
              sx={{
                display: "flex",
                flexGrow: 1,
                flexDirection: "row",
                alignItems: "center",
                justifyContent: "space-evenly",
              }}
            >
              <Typography
                variant="h4"
                color="secondary"
                component="div"
                fontWeight="bold"
                marginLeft="20px"
              >
                ARB CENTRAL
              </Typography>
              <Typography
                variant="h6"
                color="secondary"
                component="div"
                fontStyle="italic"
                marginLeft="70px"
                sx={{ display: { xs: "none", lg: "block", xl: "none" } }}
              >
                Get Available Arbitrage Oppurtunities for various bets
              </Typography>
              {isLoggedIn ? (
                <Box>
                  <Button onClick={handleClick} variant="contained">
                    <Typography variant="h6" color="secondary" component="p">
                      Log Out
                    </Typography>
                  </Button>
                </Box>
              ) : (
                <Box></Box>
              )}
            </Box>
          </Toolbar>
        </AppBar>
      </Box>
    </>
  );
};

export default Nav;
