import React, { useEffect, useState } from "react"
import { getTimeline } from "../api"
import { Card, List, Text, Grid, Button } from '@mantine/core';

export function Timeline() {
  const limit = 30;
  const [offset, setOffset] = useState(0)
  const [timeline, setTimeline] = useState<{
    Title:string;
    Description: string;
    CreatedAt:number;
  }[]>([])
  
  const getTimelineList = async () => {
    const res = await getTimeline(limit, offset)
    setTimeline(res)
    setOffset(prev => prev + limit)
  }

  useEffect(() => {
    getTimelineList()
  }, [])

  return <>
    <nav>
      <Button disabled={offset === 0}>prev</Button>
      <Button onClick={() => {
        getTimelineList()
      }}>next</Button>
    </nav>
    <List>
      {timeline.map(item => <Card key={item.CreatedAt} withBorder>
        <Text weight={700}>{item.Title}</Text>
        <Text size="sm">{item.Description}</Text>
        <Text size="xs" color="dimmed">{new Date(item.CreatedAt * 1000).toLocaleString('ja-JP')}</Text>
      </Card>)}
    </List>
  </>
}