import React from "react";
import Layout from '@theme/Layout';
import {RedocStandalone} from "redoc";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";

const ApiDoc = () => {
  const {siteConfig} = useDocusaurusContext();
  const options = {
    scrollYOffset: '.navbar',
    theme: {
      sidebar: {
        width: "300px"
      },
      typography: {
        fontSize: 'var(--ifm-font-size-base)',
        lineHeight: 'var(--ifm-line-height-base)',
        fontFamily: 'var(--ifm-font-family-base)',
        headings: {
          fontFamily: 'var(--ifm-font-family-base)',
          fontWeight: 'var(--ifm-heading-font-weight)'
        },
        code: {
          lineHeight: 'var(--ifm-pre-line-height)',
          fontFamily: 'var(--ifm-font-family-monospace)'
        }
      },
      colors: {
        primary: {
          main: "#0F3950",
        },
      }
    }
  };

  return (
    <Layout title="Api">
      <RedocStandalone specUrl="/api/specs.yaml" options={options}/>
    </Layout>
  )
}

export default ApiDoc;
