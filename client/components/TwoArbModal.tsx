import Backdrop from "@mui/material/Backdrop";
import Box from "@mui/material/Box";
import Modal from "@mui/material/Modal";
import Fade from "@mui/material/Fade";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { TextField } from "@mui/material";
import { useEffect, useState } from "react";
import { useTheme } from "@emotion/react";

type twoArbRow = {
  Away: string;
  AwayOdds: number;
  CreatedAt: string;
  DeletedAt: null;
  GameTime: string;
  GameType: string;
  Home: string;
  HomeOdds: number;
  ID: number;
  League: string;
  Profit: number;
  Title: string;
  updatedAt: string;
};

interface props {
  props: twoArbRow;
}

const TwoArbModal = ({ props }: props) => {
  const [stake, setStake] = useState(100);
  const [profit, setProfit] = useState(0);
  const [homeStake, setHomeStake] = useState({
    stake: 0,
    profit: 0,
  });
  const [awayStake, setAwayStake] = useState({
    stake: 0,
    profit: 0,
  });
  const [open, setOpen] = useState(false);
  const theme = useTheme();
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  useEffect(() => {
    const arb = 1 / props.HomeOdds + 1 / props.AwayOdds;
    //set arbProfit
    const arbProfit = stake / arb;
    setProfit(Number(arbProfit.toFixed(2)));
    //calculate stakes for each outcome
    const homeArb = (stake * 1) / props.HomeOdds / arb;
    const homeProfit = homeArb * props.HomeOdds;
    setHomeStake({
      stake: Number(homeArb.toFixed(2)),
      profit: Number(homeProfit.toFixed(2)),
    });
    const awayArb = (stake * 1) / props.AwayOdds / arb;
    const awayProfit = awayArb * props.AwayOdds;
    setAwayStake({
      stake: Number(awayArb.toFixed(2)),
      profit: Number(awayProfit.toFixed(2)),
    });
  }, [stake]);

  return (
    <div>
      <Button
        sx={{
          //@ts-ignore
          bgcolor: theme.palette.neutral.main,
        }}
        onClick={handleOpen}
      >
        <Typography
          sx={{
            //@ts-ignore
            color: theme.palette.secondary.main,
          }}
        >
          Open Arb
        </Typography>
      </Button>
      <Modal
        aria-labelledby="transition-modal-title"
        aria-describedby="transition-modal-description"
        open={open}
        onClose={handleClose}
        closeAfterTransition
        BackdropComponent={Backdrop}
        BackdropProps={{
          timeout: 500,
        }}
      >
        <Fade in={open}>
          <Box
            sx={{
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              width: 600,
              bgcolor: "background.paper",
              border: "2px solid #000",
              boxShadow: 24,
              p: 4,
            }}
          >
            <Typography id="transition-modal-title" variant="h6" component="h2">
              Arbitrage Betting Calculator
            </Typography>
            <Box>
              <Typography id="transition-modal-description" sx={{ mt: 2 }}>
                {props.Title}
              </Typography>
              <Typography id="transition-modal-description" sx={{ mt: 2 }}>
                {props.GameType.split("_")[0].toUpperCase()}
              </Typography>
              <Typography id="transition-modal-description" sx={{ mt: 2 }}>
                {props.League}
              </Typography>
              <Typography id="transition-modal-description" sx={{ mt: 2 }}>
                {props.GameTime.slice(8, 10)}/{props.GameTime.slice(5, 7)}/
                {props.GameTime.slice(0, 4)}, {props.GameTime.slice(11, 16)}
              </Typography>
            </Box>
            <Typography
              id="transition-modal-title"
              variant="h6"
              component="h2"
              marginTop="10px"
            >
              Your Stake?
            </Typography>
            <Box
              sx={{
                display: "inline-grid",
                gridAutoFlow: "row",
                gridTemplateColumns: "repeat(3, 1fr)",
                gridTemplateRows: "repeat(1, 1fr)",
                columnGap: "50px",
                justifyContent: "center",
                alignItems: "center",
              }}
            >
              <Box
                component="form"
                sx={{
                  marginTop: "20px",
                  width: "100px",
                }}
              >
                <TextField
                  type="number"
                  id="outlined-basic"
                  label="Stake"
                  variant="outlined"
                  //@ts-ignore
                  onChange={(e) => setStake(e.target.value)}
                  value={stake}
                />
              </Box>
              <Typography variant="h6"> = </Typography>
              <Typography variant="h6">{profit} </Typography>
            </Box>
            <Box
              sx={{
                marginTop: "20px",
                display: "inline-grid",
                gridAutoFlow: "row",
                gridTemplateColumns: "1.5fr repeat(5, 1fr)",
                gridTemplateRows: "repeat(1, 1fr)",
                columnGap: "20px",
                justifyContent: "center",
                alignItems: "center",
              }}
            >
              <Box
                sx={{
                  width: "200px",
                  marginTop: "15px",
                  display: "inline-grid",
                  gridAutoFlow: "column",
                  gridTemplateColumns: "repeat(1, 1fr)",
                  gridTemplateRows: "repeat(3, 1fr)",
                  rowGap: "20px",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <Typography variant="h6"> Betsite</Typography>
                <Typography variant="h6"> {props.Home}</Typography>
                <Typography variant="h6"> {props.Away}</Typography>
              </Box>
              <Box
                sx={{
                  marginTop: "15px",
                  display: "inline-grid",
                  gridAutoFlow: "column",
                  gridTemplateColumns: "repeat(1, 1fr)",
                  gridTemplateRows: "repeat(3, 1fr)",
                  rowGap: "20px",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <Typography variant="h6"> Odds</Typography>
                <Typography variant="h6">
                  {" "}
                  {props.HomeOdds.toFixed(2)}
                </Typography>
                <Typography variant="h6">
                  {" "}
                  {props.AwayOdds.toFixed(2)}
                </Typography>
              </Box>
              <Box
                sx={{
                  marginTop: "15px",
                  display: "inline-grid",
                  gridAutoFlow: "column",
                  gridTemplateColumns: "repeat(1, 1fr)",
                  gridTemplateRows: "repeat(3, 1fr)",
                  rowGap: "20px",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <Typography variant="h6"> </Typography>
                <Typography variant="h6"> * </Typography>
                <Typography variant="h6"> * </Typography>
              </Box>
              <Box
                sx={{
                  marginTop: "15px",
                  display: "inline-grid",
                  gridAutoFlow: "column",
                  gridTemplateColumns: "repeat(1, 1fr)",
                  gridTemplateRows: "repeat(3, 1fr)",
                  rowGap: "20px",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <Typography variant="h6"> Stake </Typography>
                <Typography variant="h6"> {homeStake.stake}</Typography>
                <Typography variant="h6"> {awayStake.stake}</Typography>
              </Box>
              <Box
                sx={{
                  marginTop: "15px",
                  display: "inline-grid",
                  gridAutoFlow: "column",
                  gridTemplateColumns: "repeat(1, 1fr)",
                  gridTemplateRows: "repeat(3, 1fr)",
                  rowGap: "20px",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <Typography variant="h6"> </Typography>
                <Typography variant="h6"> = </Typography>
                <Typography variant="h6"> = </Typography>
              </Box>
              <Box
                sx={{
                  marginTop: "15px",
                  display: "inline-grid",
                  gridAutoFlow: "column",
                  gridTemplateColumns: "repeat(1, 1fr)",
                  gridTemplateRows: "repeat(3, 1fr)",
                  rowGap: "20px",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <Typography variant="h6">Return </Typography>
                <Typography variant="h6">{homeStake.profit} </Typography>
                <Typography variant="h6">{awayStake.profit} </Typography>
              </Box>
            </Box>
            <Typography variant="h6" fontWeight="bold" marginTop="20px">
              {" "}
              Sure Profit
            </Typography>
            <Typography variant="h6" marginTop="10px">
              {(profit - stake).toFixed(2)}{" "}
            </Typography>
          </Box>
        </Fade>
      </Modal>
    </div>
  );
};

export default TwoArbModal;
