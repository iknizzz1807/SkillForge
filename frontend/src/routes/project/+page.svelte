<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  type ProjectDisplay = {
    id: string;
    title: string;
    description: string;
    skills: string[];
    start_time: Date;
    end_time: Date;
    created_by_id: string;
    created_by_name: string;
    status: string;
    max_member: number;
    current_member: number;
    difficulty: string;
    created_at: string;
  };

  const token = data.token;

  let projectsDisplay: ProjectDisplay[] = $state(data.projects);
  const applicationCount: number = data.applicationCount;
  let errorLoadingProjects: string | null = $state(data.error);

  let projectsDisplayCopy = projectsDisplay;

  let filterState: string = $state("all");

  // Modal states
  let showEditModal: boolean = $state(false);
  let showDeleteModal: boolean = $state(false);
  let currentProject: ProjectDisplay | null = $state(null);

  let editProjectError: string | null = $state(null);

  // Thêm state cho modal sinh viên đang apply
  let showApplicantsModal: boolean = $state(false);
  let currentApplicants: any[] = $state([]);

  // Mock data cho sinh viên đang apply
  const mockApplicants = [
    {
      id: "user123",
      name: "John Doe",
      major: "Computer Science",
      skills: ["JavaScript", "React"],
      avatar: "https://avatars.githubusercontent.com/u/123456?v=4",
      appliedDate: "2025-04-15",
    },
    {
      id: "user456",
      name: "Jane Smith",
      major: "Data Science & Analytics",
      skills: ["Python", "Data Science"],
      avatar: "https://avatars.githubusercontent.com/u/234567?v=4",
      appliedDate: "2025-04-18",
    },
    {
      id: "user789",
      name: "Robert Chen",
      major: "Software Engineering",
      skills: ["Java", "Spring Boot"],
      avatar: "https://avatars.githubusercontent.com/u/345678?v=4",
      appliedDate: "2025-04-20",
    },
  ];

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

  // Add these functions for skill selection
  function toggleSkill(skill: string) {
    if (editSkills.includes(skill)) {
      editSkills = editSkills.filter((s) => s !== skill);
    } else {
      editSkills = [...editSkills, skill];
    }
  }

  function removeSkill(skill: string) {
    editSkills = editSkills.filter((s) => s !== skill);
  }

  let skillSearchTerm: string = $state("");
  let showSkillsDropdown: boolean = $state(false);

  // Thêm state cho modal xác nhận xóa
  let showRemoveStudentModal: boolean = $state(false);
  let studentToRemove: any = $state(null);

  // Thêm state cho modal xác nhận rời dự án
  let showLeaveProjectModal: boolean = $state(false);
  let projectToLeave: ProjectDisplay | null = $state(null);

  // Hàm mở modal xác nhận rời dự án
  function openLeaveProjectModal(project: ProjectDisplay) {
    projectToLeave = project;
    showLeaveProjectModal = true;
  }

  // Hàm đóng modal xác nhận rời dự án
  function closeLeaveProjectModal() {
    showLeaveProjectModal = false;
    projectToLeave = null;
  }

  // Hàm xử lý rời dự án
  const leaveProject = async () => {
    if (!projectToLeave) return;

    try {
      // TODO: Thay thế bằng API call thực tế
      // const response = await fetch(`/api/projects/${projectToLeave.id}/leave`, {
      //   method: 'POST',
      //   headers: {
      //     'Authorization': `Bearer ${token}`
      //   }
      // });

      // if (!response.ok) {
      //   const errorData = await response.json();
      //   throw new Error(errorData.error || 'Failed to leave project');
      // }

      // Mock xử lý: Giả lập xóa dự án khỏi danh sách hiển thị hoặc thay đổi trạng thái
      projectsDisplay = projectsDisplay.filter(
        (p) => p.id !== projectToLeave?.id
      );
      projectsDisplayCopy = projectsDisplay;

      // Đóng modal xác nhận
      closeLeaveProjectModal();

      // Hiển thị thông báo thành công
      alert(`You have successfully left the project "${projectToLeave.title}"`);
    } catch (error) {
      console.error("Error leaving project:", error);
      alert("Failed to leave project: " + error);
    }
  };

  // Hàm mở modal xác nhận xóa
  function openRemoveStudentModal(e: Event, student: any, projectId: string) {
    e.preventDefault(); // Ngăn chặn sự kiện nổi bọt
    e.stopPropagation(); // Ngăn chặn click event truyền lên thẻ cha
    studentToRemove = { ...student, projectId };
    showRemoveStudentModal = true;
  }

  function closeRemoveStudentModal() {
    showRemoveStudentModal = false;
    studentToRemove = null;
  }

  // Hàm xử lý xóa sinh viên khỏi dự án
  const removeStudentFromProject = async () => {
    if (!studentToRemove) return;

    try {
      // TODO: Thay thế bằng API call thực tế
      // const response = await fetch(`/api/projects/${studentToRemove.projectId}/students/${studentToRemove.id}`, {
      //   method: 'DELETE',
      //   headers: {
      //     'Authorization': `Bearer ${token}`
      //   }
      // });

      // if (!response.ok) {
      //   const errorData = await response.json();
      //   throw new Error(errorData.error || 'Failed to remove student');
      // }

      // Xóa sinh viên khỏi danh sách hiện tại
      currentApplicants = currentApplicants.filter(
        (a) => a.id !== studentToRemove.id
      );

      // Đóng modal xác nhận
      closeRemoveStudentModal();

      // Hiển thị thông báo thành công
      alert(`Removed ${studentToRemove.name} from the project successfully`);
    } catch (error) {
      console.error("Error removing student:", error);
      alert("Failed to remove student: " + error);
    }
  };

  // Hàm mở modal các sinh viên đang apply
  function openApplicantsModal(projectId: string) {
    // TODO: Thực tế cần fetch danh sách sinh viên đang apply từ API
    // const fetchApplicants = async (projectId) => {
    //   try {
    //     const response = await fetch(`/api/projects/${projectId}/applicants`, {
    //       headers: {
    //         Authorization: `Bearer ${token}`
    //       }
    //     });
    //     if (!response.ok) throw new Error('Failed to fetch applicants');
    //     const applicants = await response.json();
    //     return applicants;
    //   } catch (error) {
    //     console.error('Error fetching applicants:', error);
    //     return [];
    //   }
    // };

    // Hiện tại dùng mock data
    currentApplicants = mockApplicants;
    showApplicantsModal = true;
  }

  function closeApplicantsModal() {
    showApplicantsModal = false;
    currentApplicants = [];
  }

  // Hàm format ngày apply
  function formatApplyDate(dateString: string): string {
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  }

  // Modal functions
  function openEditModal(project: ProjectDisplay) {
    currentProject = { ...project };
    showEditModal = true;
  }

  function openDeleteModal(project: ProjectDisplay) {
    currentProject = project;
    showDeleteModal = true;
  }

  function closeEditModal() {
    showEditModal = false;
    currentProject = null;
    editProjectError = null;
  }

  function closeDeleteModal() {
    showDeleteModal = false;
    currentProject = null;
  }

  // Edit form state
  let editTitle: string = $state("");
  let editDescription: string = $state("");
  let editSkills: string[] = $state([]);
  let editStartDate: string = $state("");
  let editEndDate: string = $state("");
  let editMaxMember: number = $state(0);
  let editStatus: string = $state("");
  let editDifficulty: string = $state("");

  $effect(() => {
    if (currentProject && showEditModal) {
      editTitle = currentProject.title;
      editDescription = currentProject.description;
      editSkills = currentProject.skills;
      editStartDate = new Date(currentProject.start_time)
        .toISOString()
        .split("T")[0];
      editEndDate = new Date(currentProject.end_time)
        .toISOString()
        .split("T")[0];
      editMaxMember = currentProject.max_member;
      editStatus = currentProject.status;
      editDifficulty = currentProject.difficulty;
    }
  });

  // Thêm một computed state để lọc projects dựa trên filterState
  $effect(() => {
    if (filterState === "all") {
      projectsDisplay = projectsDisplayCopy;
    } else if (filterState === "open") {
      projectsDisplay = projectsDisplayCopy.filter(
        (project) => project.status.toLowerCase() === "open"
      );
    } else if (filterState === "close") {
      projectsDisplay = projectsDisplayCopy.filter(
        (project) => project.status.toLowerCase() === "close"
      );
    }
  });

  // Close skills dropdown when clicking outside
  $effect(() => {
    function handleClickOutside(event: MouseEvent) {
      const target = event.target as HTMLElement;
      if (showSkillsDropdown && !target.closest(".skills-dropdown-container")) {
        showSkillsDropdown = false;
      }
    }

    document.addEventListener("click", handleClickOutside);
    return () => {
      document.removeEventListener("click", handleClickOutside);
    };
  });

  // Format date function
  function formatDate(dateString: string): string {
    if (!dateString) return "N/A";

    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "Invalid date";

    // Format: "Mar 24, 2025"
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  }

  // Updated to use our server routes
  const deleteProject = async () => {
    try {
      if (!currentProject) return;

      const response = await fetch(`/api/projects/${currentProject.id}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to delete project");
      }

      // Remove the deleted project from the display list
      projectsDisplay = projectsDisplay.filter(
        (p) => p.id !== currentProject?.id
      );
      projectsDisplayCopy = projectsDisplay;

      // Close the modal after successful deletion
      closeDeleteModal();
    } catch (error) {
      console.error("Error deleting project:", error);
      alert("Failed to delete project: " + error);
    }
  };

  // Updated to use our server routes
  const updateProject = async () => {
    try {
      editProjectError = null;
      if (!currentProject) return;

      // Get selected skills from the hidden select element
      const skillsSelect = document.getElementById(
        "edit-skills"
      ) as HTMLSelectElement;
      const selectedSkills = Array.from(skillsSelect.options)
        .filter((option) => option.selected)
        .map((option) => option.value);

      const response = await fetch(`/api/projects/${currentProject.id}`, {
        method: "PUT",
        body: JSON.stringify({
          title: editTitle,
          description: editDescription,
          skills: editSkills,
          start_time: new Date(editStartDate),
          end_time: new Date(editEndDate),
          max_member: editMaxMember,
          status: editStatus,
          difficulty: editDifficulty,
        }),
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to update project");
      }

      // Update the local state with the edited project
      const updatedProject = await response.json();
      projectsDisplay = projectsDisplay.map((p) =>
        p.id === currentProject?.id ? updatedProject : p
      );

      projectsDisplayCopy = projectsDisplay;

      // Close the modal after successful update
      closeEditModal();
    } catch (error) {
      console.error("Error updating project:", error);
      alert("Failed to update project: " + error);
    }
  };

  function calculateProgress(
    startDate: Date | string,
    endDate: Date | string
  ): number {
    const start = new Date(startDate);
    const end = new Date(endDate);
    const now = new Date();

    // Kiểm tra ngày hợp lệ
    if (isNaN(start.getTime()) || isNaN(end.getTime())) {
      return 0;
    }

    // Dự án chưa bắt đầu
    if (now < start) return 0;

    // Dự án đã kết thúc
    if (now > end) return 100;

    // Đang trong quá trình
    const totalDuration = end.getTime() - start.getTime();
    const elapsedDuration = now.getTime() - start.getTime();

    const progressPercent = Math.round((elapsedDuration / totalDuration) * 100);
    return progressPercent;
  }

  // Lấy màu dựa trên phần trăm tiến độ
  function getProgressColor(progressPercent: number): string {
    if (progressPercent < 25) return "#f97316"; // orange-500
    if (progressPercent < 50) return "#facc15"; // yellow-500
    if (progressPercent < 75) return "#84cc16"; // lime-500
    return "#22c55e"; // green-500
  }

  // Tính khoảng thời gian
  function calculateDuration(
    startDate: Date | string,
    endDate: Date | string
  ): string {
    const start = new Date(startDate);
    const end = new Date(endDate);

    // Kiểm tra ngày hợp lệ
    if (isNaN(start.getTime()) || isNaN(end.getTime())) {
      return "N/A";
    }

    // Tính số ngày
    const diffTime = Math.abs(end.getTime() - start.getTime());
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

    // Định dạng kết quả
    if (diffDays < 7) {
      return `${diffDays} day${diffDays !== 1 ? "s" : ""}`;
    } else if (diffDays < 30) {
      const weeks = Math.floor(diffDays / 7);
      return `${weeks} week${weeks !== 1 ? "s" : ""}`;
    } else if (diffDays < 365) {
      const months = Math.floor(diffDays / 30);
      return `${months} month${months !== 1 ? "s" : ""}`;
    } else {
      const years = Math.floor(diffDays / 365);
      return `${years} year${years !== 1 ? "s" : ""}`;
    }
  }
</script>

<header class="flex justify-between items-center mb-4 ml-64 pr-4 pl-4 pt-4">
  <button class="btn-secondary">
    <a href="/project/application">
      <div class="flex items-center">
        <div
          class="flex items-center justify-center w-5 h-5 bg-red-500 text-white text-xs font-medium rounded-full mr-2"
        >
          {applicationCount}
        </div>
        Applications
      </div>
    </a>
  </button>

  {#if data.role === "business"}
    <button class="btn">
      <a href="/project/create">
        <div class="flex items-center">
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
          Create Project
        </div>
      </a>
    </button>
  {/if}
</header>

<main class="flex-1 pr-4 pl-4 ml-64 pt-2">
  <div class="flex space-x-4">
    <!-- Phần danh sách dự án active -->
    <div class="w-3/5 space-y-3">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">Your Active Projects</h3>
        <div class="flex space-x-2">
          <button
            class={filterState === "all"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "all")}
          >
            All
          </button>
          <button
            class={filterState === "open"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "open")}
          >
            Open
          </button>
          <button
            class={filterState === "close"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "close")}
          >
            Close
          </button>
        </div>
      </div>
      {#if projectsDisplay.length !== 0 && !errorLoadingProjects}
        {#each projectsDisplay as ProjectDisplay}
          {@const progressPercent = calculateProgress(
            ProjectDisplay.start_time,
            ProjectDisplay.end_time
          )}

          <!-- Project Card demo -->
          <a
            href={"/project/" + ProjectDisplay.id}
            class="card p-4 block relative group hover:shadow-lg transition-all duration-200 border-l-4"
            style="border-left-color: {ProjectDisplay.difficulty === 'beginner'
              ? '#4ade80'
              : ProjectDisplay.difficulty === 'intermediate'
                ? '#facc15'
                : '#ef4444'};"
          >
            <div class="flex justify-between items-start">
              <div class="flex-1">
                <!-- Header với title và tag -->
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center">
                    <h4 class="text-base font-medium mr-2">
                      {ProjectDisplay.title}
                    </h4>
                    <span
                      class="text-xs px-2 py-0.5 rounded-full font-medium"
                      style="background-color: {ProjectDisplay.difficulty ===
                      'beginner'
                        ? '#ecfdf5'
                        : ProjectDisplay.difficulty === 'intermediate'
                          ? '#fef9c3'
                          : '#fee2e2'}; 
                color: {ProjectDisplay.difficulty === 'beginner'
                        ? '#065f46'
                        : ProjectDisplay.difficulty === 'intermediate'
                          ? '#854d0e'
                          : '#b91c1c'};"
                    >
                      {ProjectDisplay.difficulty?.charAt(0).toUpperCase() +
                        ProjectDisplay.difficulty?.slice(1) || "N/A"}
                    </span>
                  </div>

                  <!-- Status tag đưa sang bên phải -->
                  <span
                    class="text-xs px-2 py-1 rounded-full font-bold"
                    style="background-color: {ProjectDisplay.status.toLowerCase() ===
                    'open'
                      ? '#dcfce7'
                      : '#fee2e2'}; 
                color: {ProjectDisplay.status.toLowerCase() === 'open'
                      ? '#166534'
                      : '#b91c1c'};"
                  >
                    {ProjectDisplay.status.toUpperCase()}
                  </span>
                </div>

                <!-- Skills với thiết kế dạng tag -->
                <div class="flex flex-wrap gap-1 mb-3">
                  {#each ProjectDisplay.skills as skill}
                    <span
                      class="text-xs bg-gray-100 px-2 py-1 rounded-full text-gray-700"
                      >{skill}</span
                    >
                  {/each}
                </div>

                <!-- Progress bar -->
                <div class="mb-3">
                  <div class="flex justify-between text-xs text-gray-500 mb-1">
                    <span>Progress</span>
                    <span
                      class="font-medium"
                      style="color: {getProgressColor(progressPercent)};"
                    >
                      {progressPercent}%
                    </span>
                  </div>
                  <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                    <div
                      class="h-full"
                      style="width: {progressPercent}%; background-color: {getProgressColor(
                        progressPercent
                      )};"
                    ></div>
                  </div>
                </div>

                <!-- Thông tin thời gian với thiết kế nổi bật hơn -->
                <div
                  class="flex items-center text-xs text-gray-500 mb-3 space-x-3"
                >
                  <div class="flex items-center">
                    <svg
                      class="w-3.5 h-3.5 mr-1 text-gray-400"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke="currentColor"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                      ></path>
                    </svg>
                    <span title="Start date">
                      {formatDate(ProjectDisplay.start_time.toString())}
                    </span>
                  </div>
                  <div class="flex items-center">
                    <svg
                      class="w-3.5 h-3.5 mr-1 text-gray-400"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke="currentColor"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                      ></path>
                    </svg>
                    <span title="End date">
                      {formatDate(ProjectDisplay.end_time.toString())}
                    </span>
                  </div>
                  <div class="flex items-center">
                    <svg
                      class="w-3.5 h-3.5 mr-1 text-gray-400"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke="currentColor"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
                      ></path>
                    </svg>
                    <span title="Duration">
                      {calculateDuration(
                        ProjectDisplay.start_time,
                        ProjectDisplay.end_time
                      )}
                    </span>
                  </div>
                </div>

                <!-- Hiển thị thành viên (đã di chuyển xuống phía dưới timeline) -->
                <div class="flex items-center justify-between">
                  <div
                    class="flex items-center bg-gray-100 px-3 py-1.5 rounded-lg"
                  >
                    <div class="flex -space-x-2 mr-2">
                      <!-- Hiển thị avatars thành viên -->
                      <div
                        class="w-6 h-6 rounded-full border-2 border-white"
                        style="background-color: hsl({30 +
                          40 * ProjectDisplay.current_member}, 80%, 65%);"
                      >
                        <span
                          class="text-[9px] font-bold text-white flex items-center justify-center h-full"
                        >
                          {ProjectDisplay.current_member > 0
                            ? ProjectDisplay.current_member >= 4
                              ? ProjectDisplay.current_member
                              : ["", "JD", "SA", "MT"][
                                  ProjectDisplay.current_member
                                ]
                            : ""}
                        </span>
                      </div>
                      {#if ProjectDisplay.current_member >= 2}
                        <div
                          class="w-6 h-6 rounded-full bg-indigo-500 border-2 border-white flex items-center justify-center"
                        ></div>
                      {/if}
                      {#if ProjectDisplay.current_member >= 3}
                        <div
                          class="w-6 h-6 rounded-full bg-blue-500 border-2 border-white flex items-center justify-center"
                        ></div>
                      {/if}
                    </div>
                    <span class="text-xs font-medium">
                      {ProjectDisplay.current_member}/{ProjectDisplay.max_member}
                    </span>
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 24 24"
                      fill="currentColor"
                      class="w-4 h-4 ml-1 text-gray-500"
                    >
                      <path
                        d="M4.5 6.375a4.125 4.125 0 118.25 0 4.125 4.125 0 01-8.25 0zM14.25 8.625a3.375 3.375 0 116.75 0 3.375 3.375 0 01-6.75 0zM1.5 19.125a7.125 7.125 0 0114.25 0v.003l-.001.119a.75.75 0 01-.363.63 13.067 13.067 0 01-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 01-.364-.63l-.001-.122zM17.25 19.128l-.001.144a2.25 2.25 0 01-.233.96 10.088 10.088 0 005.06-1.01.75.75 0 00.42-.643 4.875 4.875 0 00-6.957-4.611 8.586 8.586 0 011.71 5.157v.003z"
                      />
                    </svg>
                  </div>

                  <!-- Arrow icon cho navigation -->
                  <svg
                    class="w-5 h-5 text-gray-400 group-hover:translate-x-1 transition-transform duration-200"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M9 5l7 7-7 7"
                    ></path>
                  </svg>
                </div>
              </div>
            </div>
          </a>
          <!-- Project action buttons - Only show for business users and only on hover -->
          {#if data.role === "business"}
            <div
              class="group-hover:opacity-100 transition-opacity flex space-x-2"
            >
              <button
                class="p-1.5 bg-blue-100 text-blue-600 rounded hover:bg-blue-200"
                onclick={() => openEditModal(ProjectDisplay)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                  />
                </svg>
              </button>

              <!-- Nút mới - Quản lý ứng viên -->
              <button
                class="p-1.5 bg-purple-100 text-purple-600 rounded hover:bg-purple-200"
                onclick={(e) => {
                  e.preventDefault(); // Ngăn chặn chuyển hướng từ thẻ a
                  openApplicantsModal(ProjectDisplay.id);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                  />
                </svg>
              </button>

              <button
                class="p-1.5 bg-red-100 text-red-600 rounded hover:bg-red-200"
                onclick={() => {
                  openDeleteModal(ProjectDisplay);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                  />
                </svg>
              </button>
            </div>
          {:else if data.role === "student"}
            <div
              class="group-hover:opacity-100 transition-opacity flex space-x-2"
            >
              <button
                class="p-1.5 bg-purple-100 text-purple-600 rounded hover:bg-purple-200"
                onclick={(e) => {
                  e.preventDefault(); // Ngăn chặn chuyển hướng từ thẻ a
                  openApplicantsModal(ProjectDisplay.id);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                  />
                </svg>
              </button>
              <button
                class="p-1.5 bg-red-100 text-red-600 rounded hover:bg-red-200"
                onclick={(e) => {
                  e.preventDefault();
                  openLeaveProjectModal(ProjectDisplay);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                  />
                </svg>
              </button>
            </div>
          {/if}
        {/each}
      {:else}
        <div class="card p-4 text-center">
          <p class="text-gray-500">No projects found.</p>
          <!-- <button class="btn mt-2">Create Your First Project</button> -->
        </div>
      {/if}
    </div>

    {#if data.role === "business"}
      <!-- Right sidebar with analytics and talent pool -->
      <div class="w-2/5 space-y-4">
        <!-- Business Analytics -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Project Analytics</h3>
          <div class="space-y-4">
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Active Projects</span>
                <span class="text-sm font-medium">4</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-[#6b48ff]" style="width: 80%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Talent Allocation</span>
                <span class="text-sm font-medium">12/15</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-[#ff6f61]" style="width: 75%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Tasks Completion</span>
                <span class="text-sm font-medium">92%</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-green-500" style="width: 92%"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Active Talent Pool -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Talent Pool</h3>
          <div class="space-y-2">
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/234567?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Sarah Johnson</p>
                  <p class="text-xs text-gray-500">React, Node.js</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/345678?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Michael Chen</p>
                  <p class="text-xs text-gray-500">Python, MongoDB</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/456789?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Alex Rivera</p>
                  <p class="text-xs text-gray-500">UI/UX, Vue.js</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div class="flex justify-center mt-2">
              <button class="text-sm text-[#6b48ff] hover:underline">
                View All Talent
              </button>
            </div>
          </div>
        </div>
      </div>
    {/if}
  </div>

  {#if showApplicantsModal}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div
        class="bg-white rounded-lg p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto"
      >
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-xl font-semibold">Project Applicants</h2>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeApplicantsModal}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>

        {#if currentApplicants.length > 0}
          <div class="space-y-4">
            {#each currentApplicants as applicant}
              <div
                class=" rounded-lg p-4 flex items-center justify-between bg-gray-100 hover:bg-gray-100 transition-colors"
              >
                <div class="flex items-center space-x-4">
                  <img
                    src={applicant.avatar}
                    alt={applicant.name}
                    class="w-12 h-12 rounded-full object-cover border-2 border-purple-200"
                  />
                  <div>
                    <h3 class="font-medium">{applicant.name}</h3>
                    <!-- Thêm hiển thị chuyên ngành -->
                    <p class="text-xs text-gray-600 mt-0.5">
                      {applicant.major}
                    </p>
                    <p class="text-sm text-gray-500">
                      Skills: {applicant.skills.join(", ")}
                    </p>
                    <p class="text-xs text-gray-400">
                      Applied: {formatApplyDate(applicant.appliedDate)}
                    </p>
                  </div>
                </div>
                <div class="flex space-x-2">
                  <!-- Nút xem hồ sơ sinh viên -->
                  <a
                    href={`/profile/${applicant.id}`}
                    class="px-3 py-1.5 bg-[#6b48ff] text-white text-sm rounded hover:bg-[#5a3dd3]"
                  >
                    View Profile
                  </a>

                  <!-- Thêm nút xóa sinh viên -->
                  <button
                    class="px-3 py-1.5 bg-red-100 text-red-600 text-sm rounded hover:bg-red-200"
                    onclick={(e) =>
                      openRemoveStudentModal(
                        e,
                        applicant,
                        currentProject?.id || ""
                      )}
                  >
                    Remove
                  </button>
                </div>
              </div>
            {/each}
          </div>
        {:else}
          <div class="text-center py-8">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-12 w-12 mx-auto text-gray-400"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"
              />
            </svg>
            <p class="mt-4 text-gray-600">No applicants yet</p>
          </div>
        {/if}
      </div>
    </div>
  {/if}

  <!-- Edit Project Modal -->
  {#if showEditModal && currentProject}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div
        class="bg-white rounded-lg w-full max-w-2xl max-h-[90vh] flex flex-col"
      >
        <!-- Header cố định -->
        <div
          class="flex justify-between items-center p-6 sticky top-0 bg-white rounded-t-lg z-10"
        >
          <h2 class="text-xl font-semibold">Edit Project</h2>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeEditModal}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>

        <!-- Nội dung có thể cuộn -->
        <div class="p-6 overflow-y-auto flex-1">
          <form class="space-y-4">
            <div>
              <label for="title" class="block text-sm font-medium text-gray-700"
                >Project Title</label
              >
              <input
                type="text"
                id="title"
                bind:value={editTitle}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>

            <div>
              <label
                for="description"
                class="block text-sm font-medium text-gray-700"
                >Description</label
              >
              <textarea
                id="description"
                bind:value={editDescription}
                rows="4"
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              ></textarea>
            </div>

            <!-- Trường difficulty -->
            <div>
              <label
                for="difficulty"
                class="block text-sm font-medium text-gray-700"
              >
                Difficulty Level
              </label>
              <select
                id="difficulty"
                bind:value={editDifficulty}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              >
                <option value="beginner">Beginner</option>
                <option value="intermediate">Intermediate</option>
                <option value="expert">Expert</option>
              </select>
            </div>

            <div>
              <label
                for="skills"
                class="block text-sm font-medium text-gray-700"
              >
                Skills
              </label>
              <div class="relative mt-1 skills-dropdown-container">
                <!-- Single input field for skills -->
                <div
                  class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm cursor-pointer"
                  onclick={() => (showSkillsDropdown = !showSkillsDropdown)}
                >
                  {#if editSkills.length === 0}
                    <span class="text-gray-500">Select skills...</span>
                  {:else}
                    <span>{editSkills.length} skills selected</span>
                  {/if}
                </div>

                <!-- Dropdown for skill selection -->
                {#if showSkillsDropdown}
                  <div
                    class="absolute z-20 w-full bg-white shadow-lg max-h-60 rounded-md mt-1 overflow-auto border border-gray-300"
                    style="position: absolute; top: 100%; left: 0;"
                  >
                    <!-- Search input ONLY in the dropdown header -->
                    <div class="sticky top-0 bg-white border-b p-2">
                      <input
                        type="text"
                        placeholder="Search skills..."
                        bind:value={skillSearchTerm}
                        class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-[#6b48ff]"
                      />
                    </div>

                    <div class="p-2">
                      {#each Object.entries(skillCategories) as [category, categorySkills]}
                        <!-- Filter skills based on search term -->
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
                                    checked={editSkills.includes(skill)}
                                    class="rounded text-[#6b48ff] focus:ring-[#6b48ff]"
                                    onclick={() => toggleSkill(skill)}
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

                <!-- Selected skills chips (displayed below the input) -->
                <div class="flex flex-wrap gap-2 mt-2">
                  {#each editSkills as skill}
                    <div
                      class="bg-[#6b48ff] bg-opacity-10 text-white px-3 py-1 rounded-full text-sm flex items-center"
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
                <select class="hidden" id="edit-skills" multiple>
                  {#each availableSkills as skill}
                    <option value={skill} selected={editSkills.includes(skill)}
                      >{skill}</option
                    >
                  {/each}
                </select>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  for="startDate"
                  class="block text-sm font-medium text-gray-700"
                  >Start Date</label
                >
                <input
                  type="date"
                  id="startDate"
                  bind:value={editStartDate}
                  class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>

              <div>
                <label
                  for="endDate"
                  class="block text-sm font-medium text-gray-700"
                  >End Date</label
                >
                <input
                  type="date"
                  id="endDate"
                  bind:value={editEndDate}
                  class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  for="maxMember"
                  class="block text-sm font-medium text-gray-700"
                  >Max Members</label
                >
                <input
                  type="number"
                  id="maxMember"
                  bind:value={editMaxMember}
                  min={currentProject.current_member}
                  class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>

              <div>
                <label
                  for="status"
                  class="block text-sm font-medium text-gray-700">Status</label
                >
                <select
                  id="status"
                  bind:value={editStatus}
                  class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                >
                  <option value="open">Open</option>
                  <option value="close">Closed</option>
                </select>
              </div>
            </div>
          </form>
        </div>

        <!-- Footer cố định với các nút hành động -->
        <div
          class="p-6 sticky bottom-0 bg-white rounded-b-lg flex flex-col space-y-2 z-10"
        >
          {#if editProjectError}
            <div class="text-red-600 text-sm mb-2">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 inline mr-1"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
                />
              </svg>
              {editProjectError}
            </div>
          {/if}
          <div class="flex justify-end space-x-2 w-full">
            <button
              type="button"
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeEditModal}
            >
              Cancel
            </button>
            <button
              type="button"
              class="px-4 py-2 bg-[#6b48ff] text-white rounded-md hover:bg-[#5a3dd3]"
              onclick={updateProject}
            >
              Save Changes
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}

  <!-- Delete Project Modal -->
  {#if showDeleteModal && currentProject}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <div class="text-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-12 w-12 mx-auto text-red-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
            />
          </svg>

          <h2 class="text-xl font-semibold mt-4">Delete Project</h2>

          <p class="mt-2 text-gray-600">
            Are you sure you want to delete "{currentProject.title}"? This
            action cannot be undone.
          </p>

          <div class="flex justify-center space-x-3 mt-6">
            <button
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeDeleteModal}
            >
              Cancel
            </button>
            <button
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
              onclick={deleteProject}
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}

  <!-- Thêm modal xác nhận xóa sinh viên -->
  {#if showRemoveStudentModal && studentToRemove}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <div class="text-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-12 w-12 mx-auto text-red-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
            />
          </svg>

          <h2 class="text-xl font-semibold mt-4">Remove Team Member</h2>

          <p class="mt-2 text-gray-600">
            Are you sure you want to remove <span class="font-medium"
              ><strong>{studentToRemove.name}</strong></span
            >
            from this project? This action <strong>cannot be undone</strong>.
          </p>

          <div class="flex justify-center space-x-3 mt-6">
            <button
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeRemoveStudentModal}
            >
              Cancel
            </button>
            <button
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
              onclick={removeStudentFromProject}
            >
              Remove
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}

  <!-- Leave Project Modal -->
  {#if showLeaveProjectModal && projectToLeave}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <div class="text-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-12 w-12 mx-auto text-red-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
            />
          </svg>

          <h2 class="text-xl font-semibold mt-4">Leave Project</h2>

          <p class="mt-2 text-gray-600">
            Are you sure you want to leave <span class="font-medium"
              ><strong>{projectToLeave.title}</strong></span
            >? This action <strong>cannot be undone</strong>.
          </p>

          <div class="flex justify-center space-x-3 mt-6">
            <button
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeLeaveProjectModal}
            >
              Cancel
            </button>
            <button
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
              onclick={leaveProject}
            >
              Leave Project
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
