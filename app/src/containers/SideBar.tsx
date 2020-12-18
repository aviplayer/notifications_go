import React from "react";
import {Stack} from "@chakra-ui/layout";
import {NotificationSelector} from "../state/app/app.types";
import {SideBarItem} from "../components/SideBarItem";

export const SideBar: React.FC = React.memo(() => {
  const notifications: NotificationSelector[] = [
    {id: 1, title: "email", description: "The future can be even brighter but a goal without a plan is just a wish"},
    {id: 2, title: "test1 ", description: "The future can be even brighter but a goal without a plan is just a wish"},
    {id: 3, title: "test 2", description: "The future can be even brighter but a goal without a plan is just a wish test2 iuouuo qje fj"},
  ];
  return (
    <>
      <Stack spacing={8}>
        {notifications.map((notification) => (
          <SideBarItem title={notification.title} description={notification.description}/>
        ))}
      </Stack>
    </>
  );
});
