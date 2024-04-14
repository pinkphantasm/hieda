import useDocusaurusContext from '@docusaurus/useDocusaurusContext'
import Layout from '@theme/Layout'

import HomepageHeader from '../components/HomepageHeader'

const Home = (): JSX.Element => {
    const { siteConfig } = useDocusaurusContext()
    return (
        <Layout
            title={`${siteConfig.tagline}`}
            description="Description will go into a meta tag in <head />"
        >
            <HomepageHeader />
        </Layout>
    )
}

export default Home
