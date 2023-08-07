"use client";

import { useState } from "react";
import { AppBar, Box, Toolbar, Typography, Button } from "@mui/material";

const Nav = () => {
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
            </Box>
          </Toolbar>
        </AppBar>
      </Box>
    </>
  );
};

export default Nav;
