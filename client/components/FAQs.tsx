"use client";

import { useState } from "react";
import {
  Box,
  Collapse,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  useTheme,
  Typography,
} from "@mui/material";
import CasinoIcon from "@mui/icons-material/Casino";
import ExpandLess from "@mui/icons-material/ExpandLess";
import ExpandMore from "@mui/icons-material/ExpandMore";

interface FAQ {
  question: string;
  answer: string;
  open: boolean;
}

type FAQs = Array<FAQ>;

const FAQs = () => {
  const theme = useTheme();
  const [FAQs, setFAQs] = useState<FAQs>([
    {
      question: "What is arbitrage betting?",
      answer:
        "Arbitrage betting, also known as arbing, is a strategy used by gamblers to take advantage of discrepancies in odds offered by different bookmakers for a particular sporting event or betting market. The goal of arbitrage betting is to guarantee a profit by placing wagers on all possible outcomes of an event at odds that ensure a positive return, regardless of the event's actual outcome.",
      open: false,
    },
    {
      question: "Is arbitrage betting legal?",
      answer:
        "Arbitrage betting itself is generally legal, but it may be restricted or prohibited by some bookmakers' terms and conditions. It's essential to check the rules of individual bookmakers before engaging in arbitrage betting.",
      open: false,
    },
    {
      question: "How do I find arbitrage opportunities?",
      answer:
        "To find arbitrage opportunities, you'll need to compare odds from multiple bookmakers for the same event. Many online tools and software are available to help identify potential arbs. This is one of these online tools",
      open: false,
    },
    {
      question: "What sports are suitable for arbitrage betting?",
      answer:
        "Arbitrage opportunities can arise in various sports, but it's more common in popular sports like football, basketball, tennis, and horse racing, where multiple bookmakers offer odds.",
      open: false,
    },
    {
      question: "Do arbitrage opportunities last long?",
      answer:
        "Arbitrage opportunities are often short-lived because bookmakers frequently adjust their odds to minimize potential losses. As a result, bettors must act quickly to take advantage of arbs.",
      open: false,
    },
    {
      question: "Is there any risk involved in arbitrage betting?",
      answer:
        "While arbitrage betting aims to be risk-free, there are still some risks involved, such as human error in calculating or placing bets, or bookmakers voiding bets due to errors.",
      open: false,
    },
    {
      question: "Can I use arbitrage betting to make a living?",
      answer:
        "Making a living solely through arbitrage betting can be challenging due to the rarity of opportunities and potential account restrictions from bookmakers. Many arbers use it as a supplementary income stream.",
      open: false,
    },
    {
      question: "Can bookmakers restrict or ban arbitrage bettors?",
      answer:
        "Yes, some bookmakers frown upon arbitrage betting and may restrict or close accounts of customers who frequently engage in arbing. They do this to protect their profits and manage their risks.",
      open: false,
    },
    {
      question: "How much capital do I need for arbitrage betting?",
      answer:
        "The amount of capital required for arbitrage betting depends on the size of the potential arb and the minimum bet requirements of the bookmakers involved. Generally, a larger bankroll allows for more significant profits.",
      open: false,
    },
    {
      question: "Are there any alternatives to traditional arbitrage betting?",
      answer:
        "Some bettors explore risk-free matched betting, where they take advantage of bookmakers' free bets and promotions to generate profits.",
      open: false,
    },
    {
      question: "What betting markets are currently offered?",
      answer:
        "We offer markets in UK with future expansion expected in the EU, Australia and USA",
      open: false,
    },
    {
      question: "How often will the odds update?",
      answer:
        "Currently, I am using the free tier of an api to get odds. In future, with additional funds, i can pay for the api and render more arbs",
      open: false,
    },
  ]);

  return (
    <Box
      sx={{
        flexGrow: 1,
        justifyContent: "center",
        alignItems: "center",
        boxSizing: "border-box",
      }}
    >
      <Typography
        variant="h1" //@ts-ignore
        color={theme.palette.neutral.main}
        fontWeight="bold"
        component="div"
      >
        Our FAQs
      </Typography>
      {FAQs.map((faq, index) => {
        return (
          <Box key={index}>
            <List
              sx={{
                width: "450px",
                //@ts-ignore
                bgcolor: theme.palette.neutral.main,
                mt: "10px",
              }}
              component="nav"
              aria-labelledby="nested-list-subheader"
            >
              <ListItemButton
                onClick={() =>
                  setFAQs((prevFAQs) => {
                    const updatedFAQs = [...prevFAQs];
                    for (let i = 0; i < updatedFAQs.length; i++) {
                      if (i === index) {
                        updatedFAQs[i].open = !updatedFAQs[i].open;
                      } else {
                        updatedFAQs[i].open = false;
                      }
                    }
                    return updatedFAQs;
                  })
                }
              >
                <ListItemText primary={faq.question} />
                {faq.open ? <ExpandLess /> : <ExpandMore />}
              </ListItemButton>
              <Collapse
                in={faq.open}
                timeout="auto"
                unmountOnExit
                sx={{
                  margin: "0px",
                  padding: "0px",
                }}
              >
                <List
                  component="p"
                  sx={{ mt: "1px" }}
                  disablePadding
                  key={faq.answer}
                >
                  <ListItemButton sx={{ pl: 1 }}>
                    <ListItemIcon>
                      <CasinoIcon />
                    </ListItemIcon>
                    <ListItemText primary={faq.answer} />
                  </ListItemButton>
                </List>
              </Collapse>
            </List>
          </Box>
        );
      })}
    </Box>
  );
};

export default FAQs;
