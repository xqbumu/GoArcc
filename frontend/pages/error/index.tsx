import Head from 'next/head'
import { useRouter, withRouter } from 'next/router'
import { Fragment, ReactElement, useEffect, useState } from 'react'
import BasicLayout from '../../components/tellUsMore/layouts/BasicLayout'

const ErrorPage = (): ReactElement => {
  const router = useRouter()
  const [message, setMessage] = useState('')

  useEffect(() => {
    if (router.query.message && typeof router.query.message === 'string')
      setMessage(router.query.message)
  }, [router])

  return (
    <Fragment>
      <Head>
        <title>Error !!</title>
        <link rel="icon" href="assets/alfred.png" />
      </Head>
      <BasicLayout heading={'Alfred has encountered with error'} component={message || ''} />
    </Fragment>
  )
}

export default withRouter(ErrorPage)
