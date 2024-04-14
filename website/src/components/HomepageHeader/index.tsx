import Link from '@docusaurus/Link'
import useDocusaurusContext from '@docusaurus/useDocusaurusContext'

import styles from './styles.module.css'
import Translate from '@docusaurus/Translate'

const HomepageHeader = (): JSX.Element => {
    const { siteConfig } = useDocusaurusContext()

    return (
        <header className={styles.header}>
            <div className={styles.headerContent}></div>
            <h1 className={styles.title}>Hieda</h1>
            <h2 className={styles.tagline}>{siteConfig.tagline}</h2>
            <nav className={styles.headerButtons}>
                <Link
                    className="button button--secondary button--lg"
                    to="/docs/intro"
                >
                    <Translate>Getting started 🚀</Translate>
                </Link>
                <Link
                    className="button button--secondary button--lg"
                    to="/docs/intro"
                >
                    <Translate>Documentation 📒</Translate>
                </Link>
            </nav>
        </header>
    )
}

export default HomepageHeader
