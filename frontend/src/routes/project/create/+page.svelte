<script lang="ts">
  import { goto } from "$app/navigation";
  import { browser } from "$app/environment";

  let member: number = $state(1);
  let description: string = $state("");
  let isSubmitting: boolean = $state(false);
  let error: string | null = $state(null);

  // Khai báo các biến để lưu trữ dữ liệu form
  let title: string = $state("");
  let skills: string[] = $state([]);
  let startDate: string = $state("");
  let endDate: string = $state("");

  // Thêm state cho modal thành công
  let showSuccessModal: boolean = $state(false);
  let createdProjectInfo: any = $state(null);

  // Thêm các biến state mới
  let isSkillDropdownOpen = $state(false);
  let skillSearchTerm = $state("");
  let filteredSkills: string[] = $state([]);

  // Lọc kỹ năng theo tìm kiếm
  // Tôi đang dùng svelte 5 nên cú pháp effect có thay đổi như thế này bạn hãy nắm ý tưởng và chỉnh sửa cho đúng nhé
  $effect(() => {
    if (skillSearchTerm) {
      filteredSkills = availableSkills.filter(
        (skill) =>
          !skillSearchTerm ||
          skill.toLowerCase().includes(skillSearchTerm.toLowerCase())
      );
    }
  });
  // Toggle skill selection
  function toggleSkill(skill: string) {
    if (skills.includes(skill)) {
      skills = skills.filter((s) => s !== skill);
    } else {
      skills = [...skills, skill];
    }
  }

  // Close dropdown when clicking outside
  function handleClickOutside(event: MouseEvent) {
    const dropdown = document.getElementById("skills-dropdown");
    if (dropdown && !dropdown.contains(event.target as Node)) {
      isSkillDropdownOpen = false;
    }
  }

  // Add event listener
  import { onMount, onDestroy } from "svelte";

  onMount(() => {
    if (browser) {
      document.addEventListener("click", handleClickOutside);
    }
  });

  onDestroy(() => {
    if (browser) {
      document.removeEventListener("click", handleClickOutside);
    }
  });

  const availableSkills = [
    // Frontend
    "JavaScript",
    "TypeScript",
    "HTML/CSS",
    "React",
    "Vue",
    "Angular",
    "Svelte",
    "Next.js",
    "Nuxt.js",
    "Remix",
    "Gatsby",
    "Preact",
    "jQuery",
    "Tailwind CSS",
    "Bootstrap",
    "Material UI",
    "Chakra UI",
    "Ant Design",
    "Styled Components",
    "SASS/SCSS",
    "Redux",
    "MobX",
    "Zustand",
    "Recoil",
    "Jotai",
    "SolidJS",
    "Alpine.js",
    "WebComponents",
    "PWA",
    "Storybook",
    "Webpack",
    "Vite",
    "Parcel",
    "Rollup",
    "ESBuild",
    "Babel",
    "ESLint",
    "Prettier",
    "Jest",
    "Vitest",
    "Testing Library",
    "Cypress",
    "Playwright",

    // Backend
    "Node.js",
    "Express",
    "NestJS",
    "Fastify",
    "Koa",
    "Python",
    "Django",
    "Flask",
    "FastAPI",
    "Java",
    "Spring",
    "Spring Boot",
    "Kotlin",
    "Ktor",
    "C#",
    ".NET Core",
    "ASP.NET",
    "PHP",
    "Laravel",
    "Symfony",
    "CodeIgniter",
    "Ruby",
    "Ruby on Rails",
    "Go",
    "Gin",
    "Echo",
    "Fiber",
    "Rust",
    "Actix",
    "GraphQL",
    "REST API",
    "gRPC",
    "WebSockets",
    "Socket.IO",
    "SignalR",
    "OAuth",
    "JWT",
    "SAML",
    "Microservices",
    "Serverless",
    "WebAssembly",

    // Database
    "SQL",
    "MySQL",
    "PostgreSQL",
    "MariaDB",
    "SQLite",
    "Oracle",
    "SQL Server",
    "MongoDB",
    "Cassandra",
    "DynamoDB",
    "Couchbase",
    "Neo4j",
    "Redis",
    "Elasticsearch",
    "Firebase",
    "Supabase",
    "CockroachDB",
    "PlanetScale",
    "InfluxDB",
    "TimescaleDB",
    "Prisma",
    "TypeORM",
    "Sequelize",
    "Mongoose",
    "Knex.js",
    "JDBC",
    "JPA",
    "Entity Framework",
    "Dapper",
    "SQLAlchemy",
    "NoSQL",

    // DevOps & Cloud
    "AWS",
    "Azure",
    "Google Cloud",
    "DigitalOcean",
    "Heroku",
    "Vercel",
    "Netlify",
    "Docker",
    "Kubernetes",
    "Terraform",
    "Pulumi",
    "CloudFormation",
    "Ansible",
    "Puppet",
    "Chef",
    "CI/CD",
    "GitHub Actions",
    "Jenkins",
    "GitLab CI",
    "CircleCI",
    "Travis CI",
    "ArgoCD",
    "Flux",
    "Prometheus",
    "Grafana",
    "ELK Stack",
    "Datadog",
    "New Relic",
    "Sentry",
    "Splunk",
    "Linux",
    "Bash/Shell",
    "PowerShell",
    "Nginx",
    "Apache",
    "Istio",
    "Linkerd",
    "Envoy",
    "Load Balancing",
    "HAProxy",
    "CDN",

    // Data Science & AI
    "Python for Data Science",
    "R",
    "Data Analysis",
    "Machine Learning",
    "Deep Learning",
    "Natural Language Processing",
    "Computer Vision",
    "TensorFlow",
    "PyTorch",
    "Keras",
    "scikit-learn",
    "Pandas",
    "NumPy",
    "SciPy",
    "Matplotlib",
    "Seaborn",
    "Plotly",
    "Jupyter",
    "CUDA",
    "Big Data",
    "Data Mining",
    "Data Visualization",
    "Statistical Analysis",
    "Reinforcement Learning",
    "Time Series Analysis",
    "MLOps",
    "Hugging Face",
    "LangChain",
    "AI Ethics",
    "Generative AI",
    "Data Engineering",
    "Apache Spark",
    "Hadoop",
    "Airflow",
    "dbt",
    "Databricks",

    // Mobile Development
    "iOS Development",
    "Android Development",
    "Swift",
    "Objective-C",
    "Kotlin for Android",
    "Java for Android",
    "React Native",
    "Flutter",
    "Xamarin",
    "Ionic",
    "Capacitor",
    "NativeScript",
    "SwiftUI",
    "Jetpack Compose",
    "Mobile UI Design",
    "AR/VR Mobile",
    "Mobile Testing",
    "App Store Optimization",
    "Google Play Console",
    "Mobile Analytics",
    "Push Notifications",
    "Mobile Security",
    "Offline Storage",
    "Mobile Performance",

    // Design & UI/UX
    "UI Design",
    "UX Design",
    "Figma",
    "Adobe XD",
    "Sketch",
    "InVision",
    "Prototyping",
    "Wireframing",
    "Design Systems",
    "Photoshop",
    "Illustrator",
    "User Research",
    "Usability Testing",
    "Information Architecture",
    "Interaction Design",
    "Visual Design",
    "Accessibility (a11y)",
    "Color Theory",
    "Typography",
    "Motion Design",

    // Blockchain & Web3
    "Blockchain",
    "Ethereum",
    "Solidity",
    "Web3.js",
    "ethers.js",
    "Smart Contracts",
    "Hardhat",
    "Truffle",
    "Foundry",
    "DeFi",
    "NFT Development",
    "Crypto Wallets",
    "IPFS",
    "Solana",
    "Rust for Blockchain",
    "Zero-knowledge Proofs",
    "Consensus Algorithms",
    "Tokenomics",
    "Layer 2 Solutions",
    "Cross-chain Development",

    // Game Development
    "Unity",
    "Unreal Engine",
    "Godot",
    "C++ for Games",
    "Game Design",
    "3D Modeling",
    "2D Art",
    "Level Design",
    "Game Physics",
    "Game AI",
    "Multiplayer Networking",
    "Console Development",
    "WebGL",
    "Shader Programming",
    "Game Testing",
    "Game Audio",
    "Game Monetization",
    "Game Analytics",

    // AR/VR/XR
    "AR Development",
    "VR Development",
    "MR/XR Development",
    "ARKit",
    "ARCore",
    "Unity XR",
    "WebXR",
    "Three.js",
    "A-Frame",
    "360° Media",
    "Spatial Computing",
    "Meta Quest Development",
    "HoloLens",
    "Apple Vision Pro",
    "OpenXR",
    "Spatial Audio",
    "3D User Interfaces",
    "Haptic Feedback",

    // Security
    "Cybersecurity",
    "Ethical Hacking",
    "Penetration Testing",
    "Cryptography",
    "Security Auditing",
    "OWASP",
    "Secure Coding",
    "Authentication",
    "Authorization",
    "Encryption",
    "Network Security",
    "Security Compliance",
    "Vulnerability Scanning",
    "Threat Modeling",
    "Security Testing",
    "SAST",
    "DAST",

    // IoT & Embedded
    "IoT Development",
    "Embedded Systems",
    "Arduino",
    "Raspberry Pi",
    "ESP32",
    "MQTT",
    "CoAP",
    "Embedded C/C++",
    "PCB Design",
    "Microcontrollers",
    "RTOS",
    "Firmware Development",
    "Bluetooth/BLE",
    "Zigbee",
    "LoRaWAN",
    "Sensors & Actuators",
    "Edge Computing",
    "IoT Cloud Platforms",
    "IoT Security",

    // Project Management & Teamwork
    "Agile",
    "Scrum",
    "Kanban",
    "Project Management",
    "Jira",
    "Confluence",
    "Git",
    "GitHub",
    "GitLab",
    "Bitbucket",
    "Code Review",
    "Technical Documentation",
    "Team Leadership",
    "Product Management",
    "Technical Writing",
    "Business Analysis",
    "Requirements Gathering",
    "Stakeholder Management",
    "Software Architecture",
    "System Design",
    "DevOps Culture",

    // Soft Skills & Others
    "Problem Solving",
    "Analytical Thinking",
    "Communication",
    "Technical Presentations",
    "Mentoring",
    "Open Source Contribution",
    "Continuous Learning",
    "Critical Thinking",
    "Time Management",
    "Technical Interviews",
  ];

  // Danh sách kỹ năng nhóm theo danh mục
  const skillCategories = {
    "Frontend Core": [
      "JavaScript",
      "TypeScript",
      "HTML/CSS",
      "React",
      "Vue",
      "Angular",
      "Svelte",
      "jQuery",
    ],
    "Frontend Frameworks": [
      "Next.js",
      "Nuxt.js",
      "Remix",
      "Gatsby",
      "Preact",
      "SolidJS",
      "Alpine.js",
    ],
    "UI & Styling": [
      "Tailwind CSS",
      "Bootstrap",
      "Material UI",
      "Chakra UI",
      "Ant Design",
      "Styled Components",
      "SASS/SCSS",
      "CSS Modules",
      "CSS-in-JS",
    ],
    "Frontend State & Data": [
      "Redux",
      "MobX",
      "Zustand",
      "Recoil",
      "Jotai",
      "XState",
      "TanStack Query (React Query)",
      "SWR",
      "Apollo Client",
    ],
    "Frontend Build & Tools": [
      "Webpack",
      "Vite",
      "Parcel",
      "Rollup",
      "ESBuild",
      "Babel",
      "ESLint",
      "Prettier",
      "PWA",
      "Storybook",
      "WebComponents",
    ],
    "Frontend Testing": [
      "Jest",
      "Vitest",
      "Testing Library",
      "Cypress",
      "Playwright",
      "Puppeteer",
      "Storybook Testing",
    ],

    "Backend Languages": [
      "Node.js",
      "Python",
      "Java",
      "Kotlin",
      "C#",
      "PHP",
      "Ruby",
      "Go",
      "Rust",
      "Scala",
      "Deno",
    ],
    "Backend Frameworks": [
      "Express",
      "NestJS",
      "Fastify",
      "Koa",
      "Django",
      "Flask",
      "FastAPI",
      "Spring",
      "Spring Boot",
      "ASP.NET",
      ".NET Core",
      "Laravel",
      "Symfony",
      "Ruby on Rails",
      "Gin",
      "Echo",
      "Fiber",
      "Actix",
    ],
    "API Development": [
      "REST API",
      "GraphQL",
      "gRPC",
      "WebSockets",
      "Socket.IO",
      "SignalR",
      "API Gateway",
      "API Documentation",
      "Swagger/OpenAPI",
      "tRPC",
      "JSON:API",
    ],
    "Backend Architecture": [
      "Microservices",
      "Serverless",
      "Monolithic",
      "Event-Driven",
      "CQRS",
      "DDD",
      "Clean Architecture",
      "Hexagonal Architecture",
      "Service-Oriented Architecture",
    ],
    "Authentication & Security": [
      "OAuth",
      "JWT",
      "SAML",
      "OpenID Connect",
      "Two-Factor Authentication",
      "Role-Based Access Control",
      "API Security",
      "Authentication Flows",
      "SSO",
    ],

    "SQL Databases": [
      "SQL",
      "MySQL",
      "PostgreSQL",
      "MariaDB",
      "SQLite",
      "Oracle",
      "SQL Server",
      "Google Cloud SQL",
      "Amazon RDS",
      "Azure SQL Database",
    ],
    "NoSQL Databases": [
      "MongoDB",
      "Cassandra",
      "DynamoDB",
      "Couchbase",
      "Neo4j",
      "Firebase Firestore",
      "Redis",
      "Elasticsearch",
      "CouchDB",
      "RavenDB",
    ],
    "Database Tools & ORM": [
      "Prisma",
      "TypeORM",
      "Sequelize",
      "Mongoose",
      "Knex.js",
      "SQLAlchemy",
      "Entity Framework",
      "Dapper",
      "JDBC",
      "JPA",
      "Hibernate",
    ],
    "Database Administration": [
      "Database Design",
      "Indexing",
      "Query Optimization",
      "Database Migration",
      "Backups & Recovery",
      "Replication",
      "Sharding",
      "Connection Pooling",
      "Database Security",
    ],
    "Modern Data Solutions": [
      "Supabase",
      "PlanetScale",
      "CockroachDB",
      "TiDB",
      "InfluxDB",
      "TimescaleDB",
      "QuestDB",
      "Snowflake",
      "BigQuery",
      "Redshift",
      "Data Warehousing",
    ],

    "DevOps Fundamentals": [
      "CI/CD",
      "Docker",
      "Kubernetes",
      "Infrastructure as Code",
      "Monitoring",
      "Logging",
      "DevSecOps",
      "GitOps",
      "Continuous Deployment",
      "Release Management",
    ],
    "Cloud Platforms - AWS": [
      "AWS",
      "EC2",
      "S3",
      "RDS",
      "Lambda",
      "ECS",
      "EKS",
      "CloudFormation",
      "DynamoDB",
      "AWS CDK",
      "AWS CloudFront",
      "AWS SQS/SNS",
    ],
    "Cloud Platforms - Others": [
      "Azure",
      "Google Cloud",
      "DigitalOcean",
      "Heroku",
      "Vercel",
      "Netlify",
      "Cloudflare",
      "IBM Cloud",
      "Oracle Cloud",
      "Linode",
      "Render",
    ],
    "Infrastructure & Configuration": [
      "Terraform",
      "Pulumi",
      "Ansible",
      "Puppet",
      "Chef",
      "CloudFormation",
      "ARM Templates",
      "Vagrant",
      "Packer",
      "Configuration Management",
    ],
    "CI/CD & Automation": [
      "GitHub Actions",
      "Jenkins",
      "GitLab CI",
      "CircleCI",
      "Travis CI",
      "ArgoCD",
      "Flux",
      "TeamCity",
      "Drone CI",
      "Tekton",
      "Spinnaker",
    ],
    "Monitoring & Observability": [
      "Prometheus",
      "Grafana",
      "ELK Stack",
      "Datadog",
      "New Relic",
      "Sentry",
      "Splunk",
      "Jaeger",
      "OpenTelemetry",
      "Logging",
      "APM",
      "Metrics",
    ],
    "DevOps Tools & Networking": [
      "Linux",
      "Bash/Shell",
      "PowerShell",
      "Nginx",
      "Apache",
      "Istio",
      "Linkerd",
      "Envoy",
      "Load Balancing",
      "HAProxy",
      "CDN",
      "DNS Management",
    ],

    "Data Science Fundamentals": [
      "Python for Data Science",
      "R",
      "Data Analysis",
      "Statistical Analysis",
      "Data Visualization",
      "Exploratory Data Analysis",
      "Data Cleaning",
      "Feature Engineering",
      "Experimental Design",
    ],
    "Machine Learning": [
      "Machine Learning",
      "Supervised Learning",
      "Unsupervised Learning",
      "Classification",
      "Regression",
      "Clustering",
      "Dimensionality Reduction",
      "Ensemble Methods",
      "Feature Selection",
    ],
    "Deep Learning": [
      "Deep Learning",
      "Neural Networks",
      "CNN",
      "RNN",
      "LSTM",
      "Transformers",
      "GAN",
      "Transfer Learning",
      "Fine-tuning",
      "Reinforcement Learning",
    ],
    "ML/AI Frameworks & Libraries": [
      "TensorFlow",
      "PyTorch",
      "Keras",
      "scikit-learn",
      "Hugging Face",
      "LangChain",
      "JAX",
      "MXNet",
      "XGBoost",
      "LightGBM",
      "ONNX",
    ],
    "Data Science Tools": [
      "Pandas",
      "NumPy",
      "SciPy",
      "Matplotlib",
      "Seaborn",
      "Plotly",
      "Jupyter",
      "Dask",
      "Spark",
      "PySpark",
      "Bokeh",
      "Streamlit",
    ],
    "AI Specializations": [
      "Natural Language Processing",
      "Computer Vision",
      "Speech Recognition",
      "Generative AI",
      "Recommendation Systems",
      "Time Series Analysis",
      "Anomaly Detection",
      "AI Ethics",
    ],
    "MLOps & Productionization": [
      "MLOps",
      "ML Pipelines",
      "Model Serving",
      "Model Monitoring",
      "Feature Stores",
      "ML Versioning",
      "ML Testing",
      "Weights & Biases",
      "MLflow",
      "Kubeflow",
      "BentoML",
    ],
    "Big Data": [
      "Big Data",
      "Hadoop",
      "Spark",
      "Kafka",
      "Hive",
      "HBase",
      "Storm",
      "Flink",
      "Beam",
      "Data Lake",
      "ETL",
      "Distributed Computing",
    ],
    "Data Engineering": [
      "Data Engineering",
      "Airflow",
      "dbt",
      "Dagster",
      "Prefect",
      "Fivetran",
      "Stitch",
      "Data Pipelines",
      "Data Warehousing",
      "Data Modeling",
      "Batch Processing",
      "Stream Processing",
    ],

    "Mobile - iOS": [
      "iOS Development",
      "Swift",
      "Objective-C",
      "SwiftUI",
      "UIKit",
      "Core Data",
      "Core Animation",
      "ARKit",
      "WidgetKit",
      "TestFlight",
      "Xcode",
      "CocoaPods",
      "Swift Package Manager",
    ],
    "Mobile - Android": [
      "Android Development",
      "Kotlin for Android",
      "Java for Android",
      "Jetpack Compose",
      "Android SDK",
      "Android Jetpack",
      "Room",
      "WorkManager",
      "Navigation Component",
      "Material Design",
      "Gradle",
    ],
    "Cross-Platform Mobile": [
      "React Native",
      "Flutter",
      "Xamarin",
      "Ionic",
      "Capacitor",
      "NativeScript",
      "Cordova",
      "PWA for Mobile",
      "KMM (Kotlin Multiplatform Mobile)",
      "Unity for Mobile",
    ],
    "Mobile DevOps & QA": [
      "Mobile CI/CD",
      "Fastlane",
      "AppCenter",
      "Mobile Testing",
      "Appium",
      "Espresso",
      "XCTest",
      "Firebase Test Lab",
      "Mobile Analytics",
      "Crash Reporting",
      "Performance Monitoring",
    ],
    "Mobile Features & Services": [
      "Push Notifications",
      "Mobile Authentication",
      "In-App Purchases",
      "Mobile Maps",
      "Geolocation",
      "Bluetooth/BLE",
      "NFC",
      "Camera & Photos",
      "Offline Storage",
      "Background Processing",
    ],

    "UI/UX Design": [
      "UI Design",
      "UX Design",
      "Interaction Design",
      "Visual Design",
      "User Research",
      "Usability Testing",
      "Information Architecture",
      "Wireframing",
      "Prototyping",
      "Design Systems",
      "Responsive Design",
    ],
    "Design Tools": [
      "Figma",
      "Adobe XD",
      "Sketch",
      "InVision",
      "Zeplin",
      "Framer",
      "Photoshop",
      "Illustrator",
      "After Effects",
      "Principle",
      "ProtoPie",
    ],
    "Web Design Specialties": [
      "Web Animation",
      "Microinteractions",
      "Typography",
      "Color Theory",
      "Accessibility (a11y)",
      "Dark Mode Design",
      "Design for Performance",
      "Mobile-First Design",
      "Landing Page Design",
    ],
    "Product Design": [
      "Product Design",
      "Design Thinking",
      "User Personas",
      "User Journeys",
      "A/B Testing",
      "Design Documentation",
      "Design QA",
      "Product Strategy",
      "Brand Design",
      "Content Design",
    ],

    "Blockchain & Crypto": [
      "Blockchain",
      "Bitcoin",
      "Ethereum",
      "Solidity",
      "Web3.js",
      "ethers.js",
      "Smart Contracts",
      "Crypto Wallets",
      "Tokenomics",
      "Consensus Algorithms",
      "Private Blockchains",
    ],
    "Web3 Development": [
      "DApp Development",
      "NFT Development",
      "DeFi",
      "DAOs",
      "Hardhat",
      "Truffle",
      "Foundry",
      "IPFS",
      "The Graph",
      "ERC Standards",
      "Smart Contract Security",
      "Gas Optimization",
    ],
    "Alternative Blockchains": [
      "Solana",
      "Rust for Blockchain",
      "Polkadot",
      "Binance Smart Chain",
      "Avalanche",
      "Polygon",
      "zkRollups",
      "Layer 2 Solutions",
      "Cross-chain Development",
      "Interoperability",
    ],

    "Game Development Core": [
      "Game Design",
      "Unity",
      "Unreal Engine",
      "Godot",
      "C++ for Games",
      "Game Physics",
      "Game AI",
      "Level Design",
      "Game Mechanics",
      "Game Optimization",
      "Game Audio",
    ],
    "Game Art & Animation": [
      "3D Modeling",
      "2D Art",
      "Game Animation",
      "Character Design",
      "Environmental Design",
      "Texturing",
      "Rigging",
      "VFX",
      "Technical Art",
      "Blender",
      "Maya",
    ],
    "Game Tech & Platforms": [
      "Console Development",
      "PC Game Development",
      "Mobile Game Development",
      "WebGL",
      "HTML5 Games",
      "Shader Programming",
      "Multiplayer Networking",
      "Game Backend Services",
      "Game Analytics",
    ],

    "AR/VR/XR": [
      "AR Development",
      "VR Development",
      "MR/XR Development",
      "ARKit",
      "ARCore",
      "Unity XR",
      "WebXR",
      "Three.js",
      "A-Frame",
      "360° Media",
      "Spatial Computing",
      "3D User Interfaces",
    ],
    "AR/VR Hardware": [
      "Meta Quest Development",
      "HoloLens",
      "Apple Vision Pro",
      "Vive",
      "Windows Mixed Reality",
      "OpenXR",
      "AR Glasses",
      "Spatial Audio",
      "Haptic Feedback",
      "Motion Tracking",
    ],

    Cybersecurity: [
      "Cybersecurity",
      "Network Security",
      "Application Security",
      "Cloud Security",
      "Security Architecture",
      "Security Compliance",
      "Threat Modeling",
      "Risk Assessment",
      "Identity Management",
      "Encryption",
    ],
    "Security Testing & Hacking": [
      "Ethical Hacking",
      "Penetration Testing",
      "Security Auditing",
      "Vulnerability Scanning",
      "SAST",
      "DAST",
      "OWASP",
      "Bug Bounty",
      "Red Team",
      "Blue Team",
      "Forensics",
    ],
    "Security Implementation": [
      "Secure Coding",
      "Authentication Implementation",
      "Authorization",
      "API Security",
      "Web Security",
      "Mobile Security",
      "DevSecOps",
      "Container Security",
      "Kubernetes Security",
      "Key Management",
    ],

    "IoT & Embedded": [
      "IoT Development",
      "Embedded Systems",
      "Arduino",
      "Raspberry Pi",
      "ESP32",
      "Microcontrollers",
      "MQTT",
      "CoAP",
      "Embedded C/C++",
      "Embedded Rust",
      "RTOS",
      "Firmware Development",
    ],
    "IoT Connectivity & Ecosystem": [
      "Bluetooth/BLE",
      "Zigbee",
      "LoRaWAN",
      "Z-Wave",
      "Thread",
      "Matter",
      "5G IoT",
      "Sensors & Actuators",
      "Edge Computing",
      "IoT Cloud Platforms",
      "IoT Analytics",
      "IoT Security",
      "Digital Twins",
    ],

    "Project & Team Management": [
      "Agile",
      "Scrum",
      "Kanban",
      "Project Management",
      "Product Management",
      "Team Leadership",
      "Technical Leadership",
      "Stakeholder Management",
      "OKRs",
      "JIRA",
      "Asana",
      "Trello",
      "ClickUp",
    ],
    "Software Craftsmanship": [
      "Clean Code",
      "SOLID Principles",
      "Design Patterns",
      "Test-Driven Development",
      "Refactoring",
      "Code Review",
      "Technical Debt Management",
      "Documentation",
      "Software Architecture",
      "System Design",
    ],
    "Version Control & Collaboration": [
      "Git",
      "GitHub",
      "GitLab",
      "Bitbucket",
      "Collaborative Coding",
      "Trunk-Based Development",
      "Git Flow",
      "GitHub Flow",
      "Conventional Commits",
      "Semantic Versioning",
    ],

    "Career & Soft Skills": [
      "Problem Solving",
      "Analytical Thinking",
      "Communication",
      "Technical Presentations",
      "Technical Writing",
      "Mentoring",
      "Open Source Contribution",
      "Continuous Learning",
      "Interviewing",
      "Career Development",
      "Technical Leadership",
    ],
  };

  async function handleSubmit(event: Event) {
    event.preventDefault();
    isSubmitting = true;
    error = null;

    try {
      // Lấy dữ liệu từ form
      const formElement = event.target as HTMLFormElement;
      const formSkills = Array.from(
        formElement.querySelector("#skills") as HTMLSelectElement
      )
        .filter((option) => (option as HTMLOptionElement).selected)
        .map((option) => (option as HTMLOptionElement).value);

      // Tạo payload cho API
      const payload = {
        title: title,
        description: description,
        skills: formSkills,
        max_member: member,
        start_time: startDate,
        end_time: endDate,
      };

      // Gửi request đến API endpoint
      const response = await fetch("/api/projects", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "Failed to create project");
      }

      // Xử lý phản hồi thành công
      const createdProject = await response.json();
      console.log("Project created:", createdProject);

      // Lưu thông tin project đã tạo và hiển thị modal thành công
      createdProjectInfo = createdProject;
      showSuccessModal = true;
    } catch (err: any) {
      error = err.message || "An unexpected error occurred";
      console.error("Error creating project:", err);
    } finally {
      isSubmitting = false;
    }
  }

  // Hàm đóng modal và chuyển hướng về trang projects
  function closeModalAndRedirect() {
    showSuccessModal = false;
    goto("/project");
  }

  // Hàm đóng modal và giữ nguyên trang để tạo project mới
  function closeModalAndStay() {
    showSuccessModal = false;
    // Reset form
    title = "";
    description = "";
    skills = [];
    startDate = "";
    endDate = "";
    member = 1;

    // Reset select element (cần phải reset DOM element vì Svelte không tự động cập nhật select multiple)
    const selectElement = document.querySelector(
      "#skills"
    ) as HTMLSelectElement;
    if (selectElement) {
      Array.from(selectElement.options).forEach((option) => {
        option.selected = false;
      });
    }
  }

  // Xử lý khi người dùng chọn skills
  function handleSkillsChange(event: Event) {
    const select = event.target as HTMLSelectElement;
    skills = Array.from(select.selectedOptions).map((option) => option.value);
  }

  // Thêm hàm tính khoảng thời gian dự án
  function calculateDuration(start: string, end: string): string {
    const startDate = new Date(start);
    const endDate = new Date(end);

    // Kiểm tra ngày hợp lệ
    if (isNaN(startDate.getTime()) || isNaN(endDate.getTime())) {
      return "Invalid dates";
    }

    // Tính số ngày
    const diffTime = Math.abs(endDate.getTime() - startDate.getTime());
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

    // Tính số tuần và số ngày
    const weeks = Math.floor(diffDays / 7);
    const days = diffDays % 7;

    // Định dạng kết quả
    if (weeks === 0) {
      return `${diffDays} day${diffDays !== 1 ? "s" : ""}`;
    } else if (days === 0) {
      return `${weeks} week${weeks !== 1 ? "s" : ""}`;
    } else {
      return `${weeks} week${weeks !== 1 ? "s" : ""} and ${days} day${days !== 1 ? "s" : ""}`;
    }
  }

  // Thêm hàm xóa skill
  function removeSkill(skillToRemove: string) {
    skills = skills.filter((skill) => skill !== skillToRemove);

    // Cập nhật lại DOM
    const selectElement = document.querySelector(
      "#skills"
    ) as HTMLSelectElement;
    if (selectElement) {
      Array.from(selectElement.options).forEach((option) => {
        option.selected = skills.includes(option.value);
      });
    }
  }
</script>

<svelte:head>
  <title>Create New Project</title>
</svelte:head>

<main class="flex-1 pr-4 pl-4 ml-64 pt-4">
  <div class="flex items-center mb-6">
    <a href="/project">
      <button class="text-gray-500 hover:text-gray-700 mr-3">
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M10 19l-7-7m0 0l7-7m-7 7h18"
          ></path>
        </svg>
      </button>
    </a>
    <h2 class="text-xl font-semibold">Create New Project</h2>
  </div>

  {#if error}
    <div
      class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4"
    >
      <p>{error}</p>
    </div>
  {/if}

  <div
    class="bg-white rounded-lg shadow-lg p-6 max-w-4xl mx-auto border border-gray-200 relative overflow-hidden"
  >
    <!-- Background pattern decorative -->
    <div
      class="absolute top-0 right-0 w-40 h-40 bg-[#6b48ff] opacity-5 rounded-full -mr-10 -mt-10"
    ></div>
    <div
      class="absolute bottom-0 left-0 w-32 h-32 bg-[#ff6f61] opacity-5 rounded-full -ml-10 -mb-10"
    ></div>

    <!-- Header with icon -->
    <div class="mb-6 flex items-center">
      <div class="bg-[#6b48ff] bg-opacity-10 p-3 rounded-full mr-4">
        <svg
          class="w-6 h-6 text-[#6b48ff]"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"
          ></path>
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-gray-800">Project Details</h3>
    </div>

    <form class="space-y-6" onsubmit={handleSubmit}>
      <!-- Basic Project Info -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label
              class="block text-sm font-medium text-gray-700 mb-1"
              for="title"
            >
              Project Title<span class="text-[#ff6f61]">*</span>
            </label>
            <div class="relative">
              <input
                type="text"
                id="title"
                bind:value={title}
                class="w-full p-3 pl-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent transition-all duration-200"
                placeholder="Enter a descriptive title"
                required
              />
              <div
                class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
              >
                <svg
                  class="w-5 h-5 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"
                  ></path>
                </svg>
              </div>
            </div>
          </div>

          <div>
            <label
              for="skills"
              class="block text-sm font-medium text-gray-700 mb-1"
            >
              Required Skills<span class="text-[#ff6f61]">*</span>
            </label>
            <div class="relative" id="skills-dropdown">
              <button
                type="button"
                class="w-full p-3 pl-10 border border-gray-300 rounded-md text-left flex justify-between items-center focus:outline-none focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent"
                onclick={() => (isSkillDropdownOpen = !isSkillDropdownOpen)}
              >
                <span class={skills.length ? "text-gray-800" : "text-gray-500"}>
                  {skills.length
                    ? `${skills.length} skills selected`
                    : "Select skills"}
                </span>
                <svg
                  class="w-5 h-5 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 9l-7 7-7-7"
                  ></path>
                </svg>
              </button>
              <div
                class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
              >
                <svg
                  class="w-5 h-5 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"
                  ></path>
                </svg>
              </div>

              <!-- Dropdown menu -->
              {#if isSkillDropdownOpen}
                <div
                  class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-y-auto"
                >
                  <!-- Search box -->
                  <div
                    class="sticky top-0 bg-white p-2 border-b border-gray-200"
                  >
                    <div class="relative">
                      <input
                        type="text"
                        placeholder="Search skills..."
                        class="w-full p-2 pl-8 border border-gray-300 rounded-md text-sm"
                        bind:value={skillSearchTerm}
                      />
                      <div
                        class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
                      >
                        <svg
                          class="w-4 h-4 text-gray-400"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                          ></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- Skills list with checkboxes -->
                  <div class="p-2">
                    {#each Object.entries(skillCategories) as [category, categorySkills]}
                      <!-- Chỉ hiển thị lĩnh vực có ít nhất một kỹ năng khớp với từ khóa tìm kiếm -->
                      {@const matchedSkills = categorySkills.filter(
                        (skill) =>
                          !skillSearchTerm ||
                          skill
                            .toLowerCase()
                            .includes(skillSearchTerm.toLowerCase())
                      )}

                      {#if !skillSearchTerm || matchedSkills.length > 0 || category
                          .toLowerCase()
                          .includes(skillSearchTerm.toLowerCase())}
                        <div class="mb-3">
                          <h4
                            class="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-1"
                          >
                            {category}
                          </h4>
                          <div class="space-y-1">
                            {#each matchedSkills as skill}
                              <label
                                class="flex items-center space-x-2 p-1 hover:bg-gray-50 rounded cursor-pointer"
                              >
                                <input
                                  type="checkbox"
                                  checked={skills.includes(skill)}
                                  class="rounded text-[#6b48ff] focus:ring-[#6b48ff]"
                                  onchange={() => toggleSkill(skill)}
                                />
                                <span class="text-sm">{skill}</span>
                              </label>
                            {/each}
                          </div>
                        </div>
                      {/if}
                    {/each}
                  </div>
                </div>
              {/if}
            </div>

            <!-- Selected skills chips -->
            <div class="flex flex-wrap gap-2 mt-2">
              {#each skills as skill}
                <div
                  class="bg-[#6b48ff] bg-opacity-10 text-[#ffffff] px-3 py-1 rounded-full text-sm flex items-center"
                >
                  {skill}
                  <button
                    type="button"
                    class="ml-1"
                    onclick={() => removeSkill(skill)}
                  >
                    <svg
                      class="w-4 h-4"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M6 18L18 6M6 6l12 12"
                      ></path>
                    </svg>
                  </button>
                </div>
              {/each}
            </div>

            <!-- Hidden select for form submission -->
            <select class="hidden" id="skills" multiple required>
              {#each availableSkills as skill}
                <option value={skill} selected={skills.includes(skill)}
                  >{skill}</option
                >
              {/each}
            </select>
          </div>

          <div>
            <label
              for="max-member"
              class="block text-sm font-medium text-gray-700 mb-1"
            >
              Team Size<span class="text-[#ff6f61]">*</span>
            </label>
            <div class="mt-2">
              <div class="flex items-center space-x-2">
                <input
                  id="max-member"
                  type="range"
                  min="1"
                  max="20"
                  bind:value={member}
                  class="w-full accent-[#6b48ff]"
                />
                <span
                  class="bg-[#6b48ff] text-white w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium"
                  >{member}</span
                >
              </div>
              <div class="flex justify-between text-xs text-gray-500 mt-1">
                <span>1 member</span>
                <span>20 members</span>
              </div>
              <p class="text-xs text-gray-500 mt-2 italic">
                {#if member === 1}
                  Solo project
                {:else if member < 5}
                  Small team (2-4 members)
                {:else if member < 10}
                  Medium team (5-9 members)
                {:else}
                  Large team (10+ members)
                {/if}
              </p>
            </div>
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              Project Timeline<span class="text-[#ff6f61]">*</span>
            </label>
            <div class="grid grid-cols-2 gap-4">
              <div class="relative">
                <label class="text-xs text-gray-500 block mb-1" for="start-date"
                  >Start Date</label
                >
                <div class="relative">
                  <input
                    type="date"
                    id="start-time"
                    bind:value={startDate}
                    class="w-full p-3 pl-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent transition-all duration-200"
                    required
                  />
                  <div
                    class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
                  >
                    <svg
                      class="w-5 h-5 text-gray-400"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                      ></path>
                    </svg>
                  </div>
                </div>
              </div>
              <div class="relative">
                <label class="text-xs text-gray-500 block mb-1" for="end-time"
                  >End Date</label
                >
                <div class="relative">
                  <input
                    type="date"
                    id="end-time"
                    bind:value={endDate}
                    class="w-full p-3 pl-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent transition-all duration-200"
                    required
                  />
                  <div
                    class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
                  >
                    <svg
                      class="w-5 h-5 text-gray-400"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                      ></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
            <!-- Hiển thị độ dài dự án -->
            {#if startDate && endDate}
              <div class="mt-2">
                <div class="text-xs text-gray-600 bg-gray-100 p-2 rounded-md">
                  Project duration: <span class="font-medium"
                    >{calculateDuration(startDate, endDate)}</span
                  >
                </div>
              </div>
            {/if}
          </div>

          <!-- External Links -->
          <div>
            <label class="block text-sm font-medium mb-1" for=""
              >External References (Optional)</label
            >
            <div class="space-y-3">
              <div class="flex space-x-2">
                <input
                  type="text"
                  placeholder="GitHub Repository URL"
                  class="flex-1 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                />
                <select
                  class="p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                >
                  <option>GitHub</option>
                  <option>Figma</option>
                  <option>Documentation</option>
                  <option>Other</option>
                </select>
              </div>
              <button
                type="button"
                class="text-sm text-[#6b48ff] flex items-center"
              >
                <svg
                  class="w-4 h-4 mr-1"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                  ></path>
                </svg>
                Add Another Link
              </button>
            </div>
          </div>

          <!-- Files & References -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1"
              >Attach Files</label
            >
            <button
              type="button"
              class="border-2 border-dashed border-gray-300 rounded-md p-6 text-center hover:border-[#6b48ff] transition-colors bg-gray-50 cursor-pointer hover:bg-gray-100"
              onclick={() => {
                const fileInput = document.getElementById("fileUpload");
                if (fileInput) fileInput.click();
              }}
            >
              <svg
                class="mx-auto h-12 w-12 text-gray-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
                ></path>
              </svg>
              <p class="mt-2 text-sm text-gray-600">
                <span class="font-medium text-[#6b48ff]">Click to upload</span> or
                drag and drop
              </p>
              <p class="mt-1 text-xs text-gray-500">
                PDF, DOC, XLS, PNG, JPG up to 10MB
              </p>
              <!-- Hiển thị file đã chọn -->
              <div class="mt-4 flex flex-wrap gap-2">
                <!-- Files sẽ hiển thị ở đây -->
              </div>
              <input type="file" class="hidden" id="fileUpload" multiple />
            </button>
          </div>
        </div>
      </div>

      <!-- Project Description -->
      <div>
        <label
          class="block text-sm font-medium text-gray-700 mb-1"
          for="description"
        >
          Project Description<span class="text-[#ff6f61]">*</span>
        </label>
        <div
          class="border border-gray-300 rounded-md overflow-hidden shadow-sm hover:shadow transition-shadow"
        >
          <!-- Toolbar với style mới -->
          <div
            class="bg-gray-50 border-b border-gray-300 p-2 flex flex-wrap gap-1"
          >
            <button type="button" class="p-1.5 hover:bg-gray-200 rounded">
              <svg
                class="w-4 h-4 text-gray-700"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path d="M13 7H7v2h6V7z" />
                <path
                  fill-rule="evenodd"
                  d="M7 2a2 2 0 00-2 2v12a2 2 0 002 2h6a2 2 0 002-2V4a2 2 0 00-2-2H7zm6 1.5v13a.5.5 0 01-.5.5h-5a.5.5 0 01-.5-.5v-13a.5.5 0 01.5-.5h5a.5.5 0 01.5.5z"
                  clip-rule="evenodd"
                />
              </svg>
            </button>
            <button
              type="button"
              class="p-1.5 hover:bg-gray-200 rounded font-bold">B</button
            >
            <button type="button" class="p-1.5 hover:bg-gray-200 rounded italic"
              >I</button
            >
            <button
              type="button"
              class="p-1.5 hover:bg-gray-200 rounded underline">U</button
            >
            <span class="border-r border-gray-300 h-6 mx-1"></span>
            <button type="button" class="p-1.5 hover:bg-gray-200 rounded">
              <svg
                class="w-4 h-4 text-gray-700"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  fill-rule="evenodd"
                  d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 5a1 1 0 011-1h6a1 1 0 110 2H4a1 1 0 01-1-1z"
                  clip-rule="evenodd"
                />
              </svg>
            </button>
          </div>

          <!-- Editor content area -->
          <textarea
            id="description"
            bind:value={description}
            class="p-4 min-h-[250px] focus:outline-none w-full resize-none bg-white font-sans"
            placeholder="Describe your project requirements, goals, and expectations in detail..."
            required
          ></textarea>

          <!-- Character counter -->
          <div
            class="bg-gray-50 border-t border-gray-300 p-2 text-xs text-gray-500 flex justify-between"
          >
            <span>{description.length} characters</span>
            <span>{description.split(/\s+/).filter(Boolean).length} words</span>
          </div>
        </div>
      </div>

      <!-- Form Actions -->
      <div class="flex justify-end space-x-3 pt-6 border-t border-gray-200">
        <button
          type="button"
          class="text-gray-700 py-2.5 px-5 rounded-md border border-gray-300 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-300 transition-colors"
          onclick={() => goto("/project")}
        >
          Cancel
        </button>

        <button
          type="submit"
          class="bg-gradient-to-r from-[#6b48ff] to-[#896DFF] text-white py-2.5 px-6 rounded-md hover:from-[#5a3de6] hover:to-[#7857e6] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#6b48ff] shadow-md hover:shadow-lg transition-all duration-200 flex items-center"
          disabled={isSubmitting}
        >
          {#if isSubmitting}
            <svg
              class="animate-spin -ml-1 mr-2 h-4 w-4 text-white"
              fill="none"
              viewBox="0 0 24 24"
            >
              <circle
                class="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                stroke-width="4"
              ></circle>
              <path
                class="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
              ></path>
            </svg>
            Creating Project...
          {:else}
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 6v6m0 0v6m0-6h6m-6 0H6"
              ></path>
            </svg>
            Create Project
          {/if}
        </button>
      </div>
    </form>
  </div>

  <!-- Modal Thành Công -->
  {#if showSuccessModal}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div
        class="bg-white rounded-lg p-8 max-w-md w-full shadow-2xl transform transition-all relative overflow-hidden"
      >
        <!-- Decoration -->
        <div
          class="absolute top-0 right-0 w-40 h-40 bg-[#6b48ff] opacity-5 rounded-full -mr-10 -mt-10"
        ></div>
        <div
          class="absolute bottom-0 left-0 w-32 h-32 bg-green-500 opacity-5 rounded-full -ml-10 -mb-10"
        ></div>

        <div class="text-center relative z-10">
          <!-- Icon Success animation -->
          <div
            class="mx-auto flex items-center justify-center h-16 w-16 rounded-full bg-green-100 mb-6 animate-pulse"
          >
            <svg
              class="h-8 w-8 text-green-600"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M5 13l4 4L19 7"
              />
            </svg>
          </div>

          <h3 class="text-xl font-bold text-gray-900 mb-3">
            Project Created Successfully!
          </h3>

          <p class="text-sm text-gray-600 mb-2">
            Your project <span class="font-semibold text-gray-800"
              >{createdProjectInfo?.title || "New Project"}</span
            > has been created and is ready to get started.
          </p>

          <!-- Project info cards -->
          <div
            class="bg-gray-50 rounded-lg p-4 mb-6 mt-4 grid grid-cols-2 gap-3"
          >
            <div class="text-left">
              <p class="text-xs text-gray-500">Start date</p>
              <p class="font-medium text-gray-800">
                {new Date(
                  createdProjectInfo?.start_time
                ).toLocaleDateString() || "N/A"}
              </p>
            </div>
            <div class="text-left">
              <p class="text-xs text-gray-500">End date</p>
              <p class="font-medium text-gray-800">
                {new Date(createdProjectInfo?.end_time).toLocaleDateString() ||
                  "N/A"}
              </p>
            </div>
            <div class="text-left">
              <p class="text-xs text-gray-500">Team size</p>
              <p class="font-medium text-gray-800">
                {createdProjectInfo?.max_member || "N/A"} members
              </p>
            </div>
            <div class="text-left">
              <p class="text-xs text-gray-500">Status</p>
              <p class="text-green-600 font-medium">Active</p>
            </div>
          </div>

          <div class="flex space-x-3 justify-center">
            <button
              class="px-5 py-2.5 bg-gradient-to-r from-[#6b48ff] to-[#896DFF] text-white rounded-md hover:from-[#5a3de6] hover:to-[#7857e6] transition-colors shadow-md hover:shadow-lg flex items-center justify-center"
              onclick={closeModalAndRedirect}
            >
              <svg
                class="w-4 h-4 mr-1.5"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 5l7 7-7 7"
                ></path>
              </svg>
              Go to Projects
            </button>
            <button
              class="px-5 py-2.5 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300 transition-colors flex items-center justify-center"
              onclick={closeModalAndStay}
            >
              <svg
                class="w-4 h-4 mr-1.5"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                ></path>
              </svg>
              Create Another
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}
</main>

<style>
  [contenteditable]:empty:before {
    content: attr(placeholder);
    color: #9ca3af;
    display: block;
  }
</style>
