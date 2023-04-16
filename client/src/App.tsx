import React, { useState } from 'react'
import { Grid } from '@mantine/core';
import './App.css'
import { Timeline } from './components/timeline'

function App() {

  return (
    <div className="App">
      <h1>Timeline</h1>
      <Grid
        justify='center'
      >
        <Grid.Col span={4}>
          <Timeline/>
        </Grid.Col>
      </Grid>
      <Timeline/>
    </div>
  )
}

export default App
